package cmd

import (
	"fmt"
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
)

func Pickq(qualities []int) string {

	picked, err := fuzzyfinder.Find(
		qualities,
		func(i int) string {
			return fmt.Sprintf("%vp", qualities[i])
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return "This quality doesn't exist"
			}
			return ""
		}))

	if err != nil {
		log.Fatalln("Couldnt pick quality, aborting...")
	}
	return fmt.Sprintf("%vp", qualities[picked])

}
