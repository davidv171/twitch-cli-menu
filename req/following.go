package req

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
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

const follow_url = "https://api.twitch.tv/kraken/streams/followed"
const get_user_url = "https://api.twitch.tv/kraken/user"
const get = "GET"

// Return list of LIVE streamers on follower list
func Live() Streams {
	url := follow_url
	reqT := get
	resp, err := Send(GenReq(&reqT, &url, nil))
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

type User struct {
	Id string `json:"_id"` // Stream name
}

// All Channels that a user follows, differs from a Stream
// Currently using many for now
type Followed struct {
	Total   int `json:"_total"`
	Follows []struct {
		CreatedAt time.Time `json:"created_at"`
		Channel   struct {
			Mature                       bool        `json:"mature"`
			Status                       string      `json:"status"`
			BroadcasterLanguage          string      `json:"broadcaster_language"`
			DisplayName                  string      `json:"display_name"`
			Game                         string      `json:"game"`
			Language                     string      `json:"language"`
			Name                         string      `json:"name"`
			CreatedAt                    time.Time   `json:"created_at"`
			UpdatedAt                    time.Time   `json:"updated_at"`
			Partner                      bool        `json:"partner"`
			Logo                         string      `json:"logo"`
			VideoBanner                  string      `json:"video_banner"`
			URL                          string      `json:"url"`
			Views                        int         `json:"views"`
			Followers                    int         `json:"followers"`
			BroadcasterType              string      `json:"broadcaster_type"`
			Description                  string      `json:"description"`
			PrivateVideo                 bool        `json:"private_video"`
			//TODO: Decide what to do with channels with private videos
			PrivacyOptionsEnabled        bool        `json:"privacy_options_enabled"`
		} `json:"channel"`
		Notifications bool `json:"notifications"`
	} `json:"follows"`
}
// Return list of all streamers user follows
func All() {
	//TODO
	//First we get the user id, then we get the follows for that channel
	//GET https://api.twitch.tv/kraken/users/<user ID>/follows/channels
	url := get_user_url
	reqT := get
	resp, err := Send(GenReq(&reqT, &url, nil))
	if err != nil {
		log.Fatalln("Couldn't send request")
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse all response")
	}
	user := User{}
	err = json.Unmarshal([]byte(respData), &user)

	url = "https://api.twitch.tv/kraken/users/" + user.Id + "/follows/channels"
	fmt.Println(url)
	resp, err = Send(GenReq(&reqT, &url, nil))
	if err != nil {
		log.Fatal("Could not send request")
	}
	if resp.StatusCode > 299 {
		log.Fatal("Could not get channels", resp.Status)
	}
	respData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse all response")
	}
	log.Println(string(respData))
	//TODO: Change into separate Struct
	all := Streams{}
	err = json.Unmarshal([]byte(respData), &all)
	if err != nil {
		log.Fatalln("Couldn't unmarshal response", string(respData))
	}

}
