package req

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type Streams struct {
	Streams []Stream `json:"streams"`
}
type Stream struct {
	Chan        Channel `json:"channel"`
	Viewers     int     `json:"viewers"`      // Amount of viewers on a Live stream
	VideoHeight int     `json:"video_height"` // Video height, ex. 720 -> 720p highest quality
}
type Channel struct {
	DisplayName     string `json:"display_name"`         // Stream name
	Game            string `json:"game"`                 // Stream currently played game
	Mature          bool   `json:"mature"`               // If stream is for mature audiences
	Title           string `json:"status"`               // Stream title
	Delay           int    `json:"delay"`                // Stream delay in seconds
	CreationDate    string `json:"created_at"`           // When the twitch account was created
	Url             string `json:"url"`                  // Stream URL, easily linkable into streamlink after
	Language        string `json:"language"`             // Language of the stream
	BroadcasterLang string `json:"broadcaster_language"` // Language of the broadcast, tends to differ in foreign coverages of esports
	Description     string `json:"description"`          // Channel description
}

var follow_url string = "https://api.twitch.tv/kraken/streams/followed"

// Return list of LIVE streamers on follower list
func Live(clientId, oauth string) Streams {

	client := &http.Client{}
	req, err := http.NewRequest("GET", follow_url, nil)
	if err != nil {
		log.Fatal("Couldn't request following...", err)
	}
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	// Twitch api doesn't accept canonicalized form...
	req.Header["Client-ID"] = []string{clientId}
	req.Header.Add("Authorization", "OAuth "+oauth)

	_, err = httputil.DumpRequest(req, false)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Couldn't get followed list, quitting", err)
	}
	resp_data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse response, quitting", err)
	}
	live := Streams{}
	err = json.Unmarshal([]byte(resp_data), &live)

	if err != nil {
		log.Fatal("Couldn't unmarshal json...", err)
	}
	return live
}

// Return list of all streamers user follows
func All() {
	//TODO

	client := &http.Client{}
	client.
}
