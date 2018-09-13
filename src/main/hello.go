package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"io/ioutil"
	"encoding/json"
	"time"
	"strings"
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
			if strings.Contains(ev.Text, "meme") == true {
				resp, err := http.Get("https://api.imgflip.com/get_memes");
				body, err := ioutil.ReadAll(resp.Body)
				//body = string(body)
				var body2 map[string] interface{}
				json.Unmarshal(body, &body2)
				fmt.Printf("bODY2 : %v", body2)
					//fmt.Printf("%v", string(body))
				memeArray := body2["data"].(map[string]interface{});
				memeArray2 := memeArray["memes"].([]interface{})
				rand.Seed(time.Now().UTC().UnixNano());
				fmt.Printf("%v", len(memeArray));
				randomNum := rand.Intn(len(memeArray2));
			  finalRes := memeArray2[randomNum].(map[string]interface{});
				finalRes2 := finalRes["url"].(string)
				if err != nil {
					fmt.Printf("%s", err);
				}
				rtm.SendMessage(rtm.NewOutgoingMessage(finalRes2, "CCSDZK8P4"));
			}
		}
	}
}
