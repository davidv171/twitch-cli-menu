package cmd

import "flag"

type Cmd struct {
	Vod     bool // Check for vods after picking streamer, lists all streamers instead of only live ones
	Quality bool // Ask for picking quality after choosing stream
	Top     bool // Instead of following channels check live channels
}

func Parse() Cmd {
	var vods bool
	var quality bool
	var top bool
	flag.BoolVar(&vods, "v", false, "Browse vods after picking a streamer")
	flag.BoolVar(&quality, "q", true, "Pick video quality after picking stream")
	flag.BoolVar(&top, "t", false, "Browse top streamers on the platform based on viewership")
	flag.Parse()

	return Cmd{
		Vod:     vods,
		Quality: quality,
		Top:     top,
	}

}
