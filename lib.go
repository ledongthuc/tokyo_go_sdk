package tokyo_go_sdk

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func Test() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	addr := "tokyo.thuc.space"
	u := url.URL{Scheme: "ws", Host: addr, Path: "/socket", RawQuery: "key=bot1&name=bot1"}

	fmt.Println("DEBUG: ", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		i := 0
		for {
			defer close(done)
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if i == 100 {
				log.Printf("recv: %s", message)
				i = 0
			} else {
				i++
			}
		}
	}()

	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	counter := 0

	for {
		select {
		case <-done:
			return
		case _ = <-ticker.C:
			// rotate := `{"e": "rotate", "data": 1}`
			// fire :=`{"e": "fire"}`
			// throttle :=`{"e": "throttle", "data": 0.5}`

			var message string
			if counter == 0 {
				message = `{"e": "rotate", "data": ` + fmt.Sprint(rand.Float64()*2*math.Pi) + `}`
			} else {
				message = `{"e": "throttle", "data": ` + fmt.Sprint(float64(counter)*0.5) + `}`
			}
			fmt.Println("Write: ", message)
			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write:", err)
				return
			}
			if counter == 3 {
				counter = 0
			} else {
				counter++
			}

		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
