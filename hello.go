package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxb-2406568984-434433071778-I7t0lE5JuVuZKFbq8TmnC5Hd")
	//api.SetDebug(true)
	user, err := api.GetUserInfo("UCN87F7EJ")
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("Name: %s", user.RealName)
	}
}
