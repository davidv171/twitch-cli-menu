package cmd

import (
	"fmt"
	"go-theatron/req"
	"go-theatron/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)
type Output struct {
    Url string
    Quality string
}
//Run through the protocol of functions to run in a sequence based on flags picked
func Protocol(command Cmd) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, fall-back to private mode")
	}
	client_id := os.Getenv("THEATRON_CLIENT_ID")
	redirect_uri := os.Getenv("THEATRON_REDIRECT_URI")

	if len(redirect_uri) < 1 || len(redirect_uri) < 1 {
		client_id = "fendbm5b5q1c2820m59sbdv9z95vs4"
		redirect_uri = "https://theatron.davidv7.xyz/"
	}

	oauth := os.Getenv("THEATRON_OAUTH_KEY")
	if len(oauth) < 1 {
		log.Fatalln("Couldn't find oauth key environment variable THEATRON_OAUTH_KEY")
	}
	following := req.Live(client_id, oauth)
	//Always pick stream for now
	// If vodmode list all following streamers
	if command.Vod {
		//TODO
	}
	pstream := Picks(following)
	output := Output {
	    Url : pstream.Chan.Url,
	}
	if command.Quality {
		// get string of max video quality, build on that
		avq := utils.Qualities(pstream.VideoHeight)
		output.Quality = Pickq(avq)
	}

	fmt.Print(output.Url + " " + output.Quality)

}
