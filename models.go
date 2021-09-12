package tokyo_go_sdk

type EventType string

const (
	// server -> client
	EventTypeState     = "state"
	EventTypeID        = "id"
	EventTypeTeamNames = "teamnames"

	// client -> server
	EventTypeRotate   = "rotate"
	EventTypeThrottle = "throttle"
	EventTypeFire     = "fire"
)

/* Events */

type GeneralEvent struct {
	Event EventType `json:"e"`
}

type StateEvent struct {
	Event EventType `json:"e"`
	Data  GameInfo  `json:"data"`
}

type GameInfo struct {
	Bounds     []float64        `json:"bounds"`
	Players    Players          `json:"players"`
	Bullets    Bullets          `json:"bullets"`
	Dead       Dead             `json:"dead"`
	Scoreboard map[string]int64 `json:"scoreboard"`
}

type CurrentUserIDEvent struct {
	Event EventType `json:"e"`
	Data  int64     `json:"data"`
}

type TeamNamesEvent struct {
	Event EventType         `json:"e"`
	Data  map[string]string `json:"data"`
}

/* Commands */

type RotateCommand struct {
	Event EventType `json:"e"`
	Data  float64   `json:"data"`
}

func NewRotateCommand(radian float64) RotateCommand {
	return RotateCommand{
		Event: EventTypeRotate,
		Data:  radian,
	}
}

func (r *RotateCommand) SetDefaultEvent() {
	if r == nil {
		return
	}
	r.Event = EventTypeRotate
}

type ThrottleCommand struct {
	Event EventType `json:"e"`
	Data  float64   `json:"data"`
}

func NewThrottleCommand(speed float64) ThrottleCommand {
	return ThrottleCommand{
		Event: EventTypeThrottle,
		Data:  speed,
	}
}

func (r *ThrottleCommand) SetDefaultEvent() {
	if r == nil {
		return
	}
	r.Event = EventTypeThrottle
}

type FireCommand struct {
	Event EventType `json:"e"`
}

func (r *FireCommand) SetDefaultEvent() {
	if r == nil {
		return
	}
	r.Event = EventTypeFire
}

func NewFireCommand() FireCommand {
	return FireCommand{
		Event: EventTypeFire,
	}
}

/* Shared model */

type Players []Player

type Player struct {
	ID       int64   `json:"id"`
	Angle    float64 `json:"angle"`
	Throttle float64 `json:"throttle"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type Bullets []Bullet

type Bullet struct {
	ID       int64   `json:"id"`
	PlayerID int64   `json:"player_id"`
	Angle    float64 `json:"angle"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type Dead []ADead

type ADead struct {
	Respawn struct {
		Seconds     int64 `json:"secs_since_epoch"`
		NanoSeconds int64 `json:"nanos_since_epoch"`
	} `json:"respawn"`
	Player Player `json:"player"`
}
