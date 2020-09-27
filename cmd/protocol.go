package cmd

import (
	"fmt"
	"twitch-cli-menu/req"
	"twitch-cli-menu/utils"
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
	var following req.LiveStreams
	var output Output

	if command.Vod {
		// Pick streamers
		all := req.All()
		picked := PickAll(all)
		// Pick available VODS for streamers
		vods := req.AllVods(picked)

		vod := PickVods(vods)
		fmt.Println(vod)

	} else {
		following = req.Live()

		pstream := PickLive(following)
		output = Output{
			Url: pstream.Chan.Url,
		}

		// TODO: Add manually specifying quality, or not at all
		if command.Quality {
			// get string of max video quality, build on that
			avq := utils.Qualities(pstream.VideoHeight)
			output.Quality = Pickq(avq)

			fmt.Print(output.Url + " " + output.Quality)
		} else {
			fmt.Print(output.Url)
		}
	}

}
