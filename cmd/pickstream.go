package cmd

import (
	"fmt"
	"go-theatron/req"
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
)

func Picks(live req.Streams) req.Stream {

	picked, err := fuzzyfinder.Find(
		live.Streams,
		func(i int) string {
			return live.Streams[i].Chan.DisplayName
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return "Could not find any streams"
			}
			return fmt.Sprintf("%s\nGame %s \nViewers:%v \nMax video quality: %vp, \nIs mature: %v \nTitle: %s, \nLanguage: %s, \nDescription: %s",
				live.Streams[i].Chan.DisplayName,
				live.Streams[i].Chan.Game,
				live.Streams[i].Viewers,
				live.Streams[i].VideoHeight,
				live.Streams[i].Chan.Mature,
				live.Streams[i].Chan.Title,
				live.Streams[i].Chan.Language,
				live.Streams[i].Chan.Description)
		}))

	if err != nil {
		log.Fatalln("Couldn't initialize picker", err)
	}

	return live.Streams[picked]

}
