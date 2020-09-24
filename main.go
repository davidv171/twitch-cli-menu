package main

import (
	"twitch-cli-menu/cmd"
)

func main() {

	parsed := cmd.Parse()
	cmd.Protocol(parsed)
}
