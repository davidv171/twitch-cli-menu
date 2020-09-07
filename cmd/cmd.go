package cmd

import "flag"

type Cmd struct {
	Vod     bool // Check for vods after picking streamer
	Quality bool // Ask for picking quality after choosing stream
}

func Parse() Cmd {

	var vods bool
	var quality bool
	flag.BoolVar(&vods, "v", false, "Browse vods after picking a streamer")
	flag.BoolVar(&quality, "q", true, "Pick video quality after picking stream")
	flag.Parse()
	return Cmd{
		Vod:     vods,
		Quality: quality,
	}

}
