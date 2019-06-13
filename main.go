package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	doRefresh()
}

func doRefresh() {
	ChatRefresh()
	nexTime := time.Now().Truncate(time.Second)
	nexTime = nexTime.Add(30 * time.Second)
	time.Sleep(time.Until(nexTime))
	doRefresh()
}

func ChatRefresh() {

	type Chat struct {
		Links    string `json:"_links"`
		Count    string `json:"chatter_count"`
		Chatters struct {
			Broadcaster []string `json:"broadcaster"`
			Vips        []string `json:"vips"`
			Moderators  []string `json:"moderators"`
			Staff       []string `json:"staff"`
			Admins      []string `json:"admins"`
			GlobalMods  []string `json:"global_mods"`
			Viewers     []string `json:"viewers"`
		} `json:"chatters"`
	}

	var url string = "http://tmi.twitch.tv/group/user/calyscope/chatters"

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	var chat Chat
	json.Unmarshal(body, &chat)
	defer resp.Body.Close()
	fmt.Println("Chat debug", chat.Chatters.Viewers)
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("erreur lors de la creation du fichier", err)
	}
	for i, chatter := range chat.Chatters.Moderators {
		if chatter == "moobot" || chatter == "wizebot" {
			fmt.Println("Not allowed to be listed", chatter)
		} else {
			fmt.Println(chatter)
			file.WriteString(chatter + "\n")
		}
		i++
	}
	for i, chatter := range chat.Chatters.Viewers {
		fmt.Println(chatter)
		file.WriteString(chatter + "\n")
		i++
	}
	file.Close()
}
