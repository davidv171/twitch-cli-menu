package req

import (
	"twitch-cli-menu/utils"
	"io"
	"log"
	"net/http"
)

// Global http client, that we recycle
var (
	C *http.Client
)

func GenClient() {

	C = &http.Client{}
}

// Generate request for the client
func GenReq(reqType, url *string, reader *io.Reader) *http.Request {

	conf := utils.GetEnv()
	var req *http.Request
	va
	var err error
	if reader != nil {
		req, err = http.NewRequest(*reqType, *url, *reader)
	} else {
		req, err = http.NewRequest(*reqType, *url, nil)
	}

	if err != nil {
		log.Fatal("Couldn't create request ...", err)
	}
	// These headers are added to every request in program
	// TODO: Support non-authenticated requests fully, removing the auth portion
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	// Twitch api doesn't accept canonicalized form...
	req.Header["Client-ID"] = []string{*&conf.Cid}
	req.Header.Add("Authorization", "OAuth "+*&conf.OAuth)

	return req

}

// Send request, receive response
func Send(req *http.Request) (*http.Response, error) {

	if C == nil {
		GenClient()
	}
	resp, err := C.Do(req)
	return resp, err
}
