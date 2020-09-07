package main

import (
	//    fzf "github.com/ktr0731/go-fuzzyfinder"

	"go-theatron/cmd"
)

func main() {

	parsed := cmd.Parse()
	cmd.Protocol(parsed)
}
