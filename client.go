package tokyo_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type Client struct {
	// connection infromation
	server    string
	userKey   string
	userName  string
	conn      *websocket.Conn
	ConnReady bool

	// communication channels
	done             chan struct{}
	stateChan        chan StateEvent
	currenUserIDChan chan CurrentUserIDEvent
	teamNamesEvent   chan TeamNamesEvent

	stateEventHandler         func(StateEvent)
	currentUserIDEventHandler func(CurrentUserIDEvent)
	teamNamesEventHandler     func(TeamNamesEvent)

	// game state, feel free to access
	CurrentPlayerID *int64
	CurrentPlayer   Player
	CurrentGameInfo GameInfo
	TeamNames       map[string]string // id - name
}

func NewClient(server, userKey, userName string) *Client {
	return NewClientWithScheme(server, userKey, userName)
}

func NewClientWithScheme(server, userKey, userName string) *Client {
	return &Client{
		server:   server,
		userKey:  userKey,
		userName: userName,

		done:             make(chan struct{}, 1),
		stateChan:        make(chan StateEvent, 10),
		currenUserIDChan: make(chan CurrentUserIDEvent, 10),
		teamNamesEvent:   make(chan TeamNamesEvent, 10),
	}
}

func (c *Client) RegisterStateEventHandler(handler func(StateEvent)) {
	c.stateEventHandler = handler
}

func (c *Client) RegisterCurrentUserIDEventHandler(handler func(CurrentUserIDEvent)) {
	c.currentUserIDEventHandler = handler
}

func (c *Client) RegisterTeamNamesEventHandler(handler func(TeamNamesEvent)) {
	c.teamNamesEventHandler = handler
}

func (c *Client) Listen() error {
	if c.conn != nil {
		return errors.New("Connection's establed")
	}

	url, err := url.Parse(fmt.Sprintf("%s?key=%s&name=%s", c.server, c.userKey, c.userName))
	fmt.Println(url.String())
	if err != nil {
		return err
	}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return err
	}
	c.conn = conn
	c.ConnReady = true

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-c.done:
			return nil
		case <-interrupt:
			if err := c.Close(); err != nil {
				return err
			}
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if closeErr := c.Close(); closeErr != nil {
					return fmt.Errorf("Close connection got error %v when connection drop because of %v", closeErr, err)
				}
				return err
			}
			c.handleEvent(message)
		}
	}
	return nil
}

func (c *Client) Close() error {
	c.ConnReady = false
	if c.conn == nil {
		return errors.New("Connection is not connected yet")
	}
	err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	c.done <- struct{}{} // notify other background tasks
	if err := c.conn.Close(); err != nil {
		return err
	}
	c.conn = nil
	close(c.done)
	return nil
}

func (c *Client) Rotate(radian float64) error {
	request := NewRotateCommand(radian)
	return c.sendCommand(request)
}

func (c *Client) Throttle(speed float64) error {
	request := NewThrottleCommand(speed)
	return c.sendCommand(request)
}

func (c *Client) Fire() error {
	request := NewFireCommand()
	return c.sendCommand(request)
}

func (c *Client) sendCommand(request interface{}) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	if c.conn == nil {
		return errors.New("Connection is closed")
	}
	return c.conn.WriteMessage(websocket.TextMessage, b)
}

func (c *Client) handleEvent(rawMessage []byte) {
	var e GeneralEvent
	err := json.Unmarshal(rawMessage, &e)
	if err != nil {
		return
	}
	switch e.Event {
	case EventTypeState:
		var fullE StateEvent
		if err := json.Unmarshal(rawMessage, &fullE); err != nil {
			return
		}
		c.handleStateEvent(fullE)
	case EventTypeID:
		var fullE CurrentUserIDEvent
		if err := json.Unmarshal(rawMessage, &fullE); err != nil {
			return
		}
		c.CurrentPlayerID = &fullE.Data
		if c.stateEventHandler != nil {
			c.currentUserIDEventHandler(fullE)
		}
	case EventTypeTeamNames:
		var fullE TeamNamesEvent
		if err := json.Unmarshal(rawMessage, &fullE); err != nil {
			return
		}
		c.TeamNames = fullE.Data
		if c.stateEventHandler != nil {
			c.teamNamesEventHandler(fullE)
		}
	}
}

func (c *Client) handleStateEvent(e StateEvent) {
	c.CurrentGameInfo = e.Data
	if c.CurrentPlayerID != nil {
		for _, player := range c.CurrentGameInfo.Players {
			if player.ID == *c.CurrentPlayerID {
				c.CurrentPlayer = player
				break
			}
		}
	}

	if c.stateEventHandler != nil {
		c.stateEventHandler(e)
	}
}
