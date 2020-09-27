package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Return possible video heights, based on maximum height given
// We hardcode the heights because we KISS
func Qualities(maxq int) []int {
	// All possible qualities, twitch doesn't yet support 1440p and 2160, we future proof
	allq := []int{160, 360, 480, 720, 1080, 1440, 2160}
	//Available qualities, to offer the user
	avq := make([]int, 0)
	for _, v := range allq {
		if maxq >= v {
			avq = append(avq, v)
		}
	}
	return avq
}

// Global configuration used in every http request
type conf struct {
	Cid      string
	Redirect string
	OAuth    string
}

var (
	c *conf
)

func EnvLoad() {

	c = new(conf)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, fall-back to private mode")
	}
	c.Cid = os.Getenv("THEATRON_CLIENT_ID")
	c.Redirect = os.Getenv("THEATRON_REDIRECT_URI")

	// Hardcode in case it's missing
	if len(c.Redirect) < 1 || len(c.Redirect) < 1 {
		c.Cid = "8j7747kd5dhhyi8geqeq5wb6y1jr22"
		c.Redirect = "https://theatron.davidv7.xyz/"
	}

	c.OAuth = os.Getenv("THEATRON_OAUTH_KEY")
	if len(c.OAuth) < 1 {
		log.Fatalln("Couldn't find oauth key environment variable THEATRON_OAUTH_KEY")
	}

}

func GetEnv() *conf {
	return c

}
