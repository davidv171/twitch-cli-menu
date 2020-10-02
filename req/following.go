package req

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

type LiveStreams struct {
	Streams []LiveStream `json:"streams"`
}
type LiveStream struct {
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
// TODO: Sort by viewer numbers
func Live() LiveStreams {
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
	live := LiveStreams{}
	err = json.Unmarshal([]byte(resp_data), &live)
	sort.Slice(live.Streams, func(i, j int) bool {
		return live.Streams[i].Viewers < live.Streams[j].Viewers
	})

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
type AllFollowed struct {
	Total   int       `json:"_total"`
	Follows []Follows `json:"follows"`
}
type Follows struct {
	CreatedAt         time.Time         `json:"created_at"`
	AllFollowsChannel AllFollowsChannel `json:"channel"`
	Notifications     bool              `json:"notifications"`
}

type AllFollowsChannel struct {
	Mature              bool      `json:"mature"`
	Status              string    `json:"status"`
	BroadcasterLanguage string    `json:"broadcaster_language"`
	DisplayName         string    `json:"display_name"`
	Game                string    `json:"game"`
	Language            string    `json:"language"`
	ID                  string    `json:"_id"`
	Name                string    `json:"name"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Partner             bool      `json:"partner"`
	Logo                string    `json:"logo"`
	VideoBanner         string    `json:"video_banner"`
	URL                 string    `json:"url"`
	Views               int       `json:"views"`
	Followers           int       `json:"followers"`
	BroadcasterType     string    `json:"broadcaster_type"`
	Description         string    `json:"description"`
	PrivateVideo        bool      `json:"private_video"`
	//TODO: Decide what to do with channels with private videos
	PrivacyOptionsEnabled bool `json:"privacy_options_enabled"`
}

// Return list of all streamers user follows
func All() AllFollowed {
	//First we get the user id, then we get the follows for that channel
	//GET https://api.twitch.tv/kraken/users/<user ID>/follows/channels
	user := getUser()

	// FIXME: Get maximum amount of channels, which is 100, defaults to 25
	// TODO: Very important to only list channels with VODS, otherwise this is
	// mostly useless, can we do this without firing a request for each streamer?
	url := "https://api.twitch.tv/kraken/users/" + user.Id + "/follows/channels?limit=100&stream_type=all"
	reqT := get
	req := GenReq(&reqT, &url, nil)

	respData := sendReq(req)

	all := AllFollowed{}
	err := json.Unmarshal([]byte(respData), &all)
	if err != nil {
		log.Fatalln("Couldn't unmarshal response", string(respData))
	}
	return all

}

type AllChannelVods struct {
	ChannelVideos ChannelVideos `json:"data"`
}
type ChannelVideos []struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	PublishedAt  time.Time `json:"published_at"`
	URL          string    `json:"url"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Viewable     string    `json:"viewable"`
	ViewCount    int       `json:"view_count"`
	Language     string    `json:"language"`
	Type         string    `json:"type"`
	Duration     string    `json:"duration"`
}

// Get user ID of user, then get user videos
// TODO: Let user pick video category? -> undecided, maybe useless feature
// TODO: Allow collection browsing? -> undecided
func AllVods(channel AllFollowsChannel) ChannelVideos {

	//TODO: Fix this stupid way of setting up types
	reqT := "GET"
	// Get first 100 videos, currently of all types, archive(auto vods)
	url := "https://api.twitch.tv/helix/videos?user_id=" + channel.ID + "&first=100"
	// For some reason this API call requires a completely different header
	// TODO: Join GenBearerReq and GenReq methods, maybe pass flag?
	req := GenBearerReq(&reqT, &url, nil)
	respData := sendReq(req)

	vods := AllChannelVods{}

	err := json.Unmarshal([]byte(respData), &vods)
	if err != nil {
		log.Fatalln("Couldn't unmarshal response", string(respData))
	}
	// Filter non viewable VODs out!
	// TODO: Filtering should account for subscription status...
	// TODO: Decide if we should be filtering anyways, let users decide if they're subscribed, or put in a show-all flag?
	filteredVods := ChannelVideos{}
	for _, vod := range vods.ChannelVideos {
		if vod.Viewable == "public" {
			filteredVods = append(filteredVods, vod)
		}
	}

	return filteredVods
}

// Wrapper for request sending, we fail on any errors as that breaks workflow completely
func sendReq(req *http.Request) []byte {
	resp, err := Send(req)

	if err != nil {
		log.Fatal("Could not send request")
	}
	if resp.StatusCode > 299 {
		msg := make([]byte, 250)
		resp.Body.Read(msg)
		log.Fatal("Could not get data ", resp.Status, string(msg))
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse the response")
	}
	return respData
}

// Translate CURRENTLY authenticated user into userId, which is what's needed for some direct queries
// Fail for any parse or request error
// Use translate to get a userId of another user
func getUser() User {

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

	if err != nil {
		log.Fatalln("Could not generate response data", err)
	}
	return user
}

// Return user id of an input username
func translateUser(username string) User {

	return User{}
}


// Get top streamers on the platform
// TODO: Support vod mode for those too
// TODO: Support game picking
// TODO: Support language picking
func Top() LiveStreams {

	// Use old API because new one kinda useless for this
	url := "https://api.twitch.tv/kraken/streams/"
	reqT := get
	// Unauthenticated call
	req := GenUnauthReq(&reqT, &url, nil)

	respData := sendReq(req)

	top := LiveStreams{}
	err := json.Unmarshal([]byte(respData), &top)

	if err != nil {
		log.Fatalln("Could not generate top streamer response data ", err)
	}
	return top
}
