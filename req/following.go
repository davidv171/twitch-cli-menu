package req

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/ktr0731/go-fuzzyfinder"
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

func Following(clientId, oauth string) string {

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
	return live.Streams[pick(live)].Chan.Url
}

func pick(live Streams) int {

	picked, err := fuzzyfinder.Find(
		live.Streams,
		func(i int) string {
			return live.Streams[i].Chan.DisplayName
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return "Could not find any streams"
			}
			return fmt.Sprintf("%s\nGame %s \nViewers:%v \nMax video quality: %vp, \nIs mature: %v \nTitle: %s, \nLanguage: %s, \nDescription: %s",
				live.Streams[i].Chan.DisplayName,
				live.Streams[i].Chan.Game,
				live.Streams[i].Viewers,
				live.Streams[i].VideoHeight,
				live.Streams[i].Chan.Mature,
				live.Streams[i].Chan.Title,
				live.Streams[i].Chan.Language,
				live.Streams[i].Chan.Description)
		}))

	if err != nil {
		log.Fatalln("Couldn't initialize picker", err)
	}

	return picked

}
