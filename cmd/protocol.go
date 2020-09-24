package cmd

import (
	"fmt"
	"go-theatron/req"
	"go-theatron/utils"
)

type Output struct {
	Url     string
	Quality string
}

//Run through the protocol of functions to run in a sequence based on flags picked
func Protocol(command Cmd) {

	utils.EnvLoad()
	//Always pick stream for now
	// If vodmode list all following streamers
	var following req.Streams
	if command.Vod {
		//TODO
		req.All()
	} else {
		following = req.Live()
	}

	pstream := Picks(following)

	output := Output{
		Url: pstream.Chan.Url,
	}

	if command.Quality {
		// get string of max video quality, build on that
		avq := utils.Qualities(pstream.VideoHeight)
		output.Quality = Pickq(avq)
	}

	fmt.Print(output.Url + " " + output.Quality)

}
