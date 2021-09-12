// go run main.go -server ws://{host}/socket -key={unique id} -name={display name}
package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"time"

	tokyo "github.com/ledongthuc/tokyo_go_sdk"
)

var server = flag.String("server", "", "server host")
var userKey = flag.String("key", "", "user's key")
var userName = flag.String("name", "", "user's name")

func main() {
	flag.Parse()
	validateParams()
	log.Printf("Start server: %s, key: %s, name: %s", *server, *userKey, *userName)

	client := tokyo.NewClient(*server, *userKey, *userName)
	client.RegisterStateEventHandler(func(e tokyo.StateEvent) {
		log.Printf("State Event: %+v", e)
	})
	client.RegisterCurrentUserIDEventHandler(func(e tokyo.CurrentUserIDEvent) {
		log.Printf("User ID Event: %+v", e)
	})
	client.RegisterTeamNamesEventHandler(func(e tokyo.TeamNamesEvent) {
		log.Printf("Team names: %+v", e)
	})
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()
		for {
			_ = <-ticker.C
			if !client.ConnReady {
				continue
			}
			client.Rotate(rand.Float64() * 2 * math.Pi)
			client.Throttle(1)
			client.Fire()
		}
	}()
	log.Fatal(client.Listen())
}

func validateParams() {
	if server == nil {
		panic("miss server flag")
	}
	if userKey == nil {
		panic("miss key flag")
	}
	if userName == nil {
		panic("miss name flag")
	}
}
