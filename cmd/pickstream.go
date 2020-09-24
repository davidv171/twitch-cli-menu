package cmd

import (
	"fmt"
	"log"
	"twitch-cli-menu/req"

	"github.com/ktr0731/go-fuzzyfinder"
)

func PickLive(live req.LiveStreams) req.LiveStream {

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

func PickAll(all req.AllFollowed) req.AllFollowsChannel {

	picked, err := fuzzyfinder.Find(
		all.Follows,
		func(i int) string {
			return all.Follows[i].AllFollowsChannel.DisplayName
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return "Could not find any streams"
			}
			return fmt.Sprintf("%s\nFollowers: %v,\nStatus: %s,\nStatus: %s,\nUpdatedAt: %s,\nMature?: %v,\nLanguage: %s",
				all.Follows[i].AllFollowsChannel.DisplayName,
				all.Follows[i].AllFollowsChannel.Followers,
				all.Follows[i].AllFollowsChannel.Status,
				all.Follows[i].AllFollowsChannel.UpdatedAt,
				all.Follows[i].AllFollowsChannel.Mature,
				all.Follows[i].AllFollowsChannel.BroadcasterLanguage)

		}))

	if err != nil {
		log.Fatalln("Couldn't initialize picker", err)
	}

	return all.Follows[picked].AllFollowsChannel

}
