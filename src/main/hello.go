package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxb-2406568984-434433071778-I7t0lE5JuVuZKFbq8TmnC5Hd")
	//api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Println("Event received")

		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
		case *slack.MessageEvent:
			fmt.Printf("Got message: %v", ev)
		}
	}
}
