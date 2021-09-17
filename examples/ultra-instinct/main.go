// go run main.go -server ws://{host}/socket -key={unique id} -name={display name}
package main

import (
	"flag"
	"log"
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
	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)
		defer ticker.Stop()

		fireCounter := 0
		for {
			_ = <-ticker.C
			if !client.ConnReady {
				continue
			}

			if fireCounter == 6 {
				otherPlayer, _, err := client.GetClosestPlayer()
				if err == nil {
					client.HeadToPoint(otherPlayer.X, otherPlayer.Y)
					client.Fire()
				}
				fireCounter = 0
			}

			if dodgeAngle, yes := client.AnyBulletToMe(500); yes {
				client.Throttle(1)
				client.Rotate(dodgeAngle)
			} else if dodgeAngle, yes := client.AnyPlayersToMe(500); yes {
				client.Throttle(1)
				client.Rotate(dodgeAngle)
			} else {
				client.Throttle(0)
			}

			fireCounter++

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
