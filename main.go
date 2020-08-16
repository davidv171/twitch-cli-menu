package main

import (
	//    fzf "github.com/ktr0731/go-fuzzyfinder"
	"fmt"
	"github.com/joho/godotenv"
	"go-theatron/req"
	"os"
)

type Stream struct {
	Name    string
	Game    string
	Viewers uint32
	Id      string
	Quality string
}
type Category struct {
	Name    string
	Viewers uint32
	Streams []Stream
}

func main() {
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
	stream := req.Following(client_id, oauth)
	fmt.Println(stream)

}
