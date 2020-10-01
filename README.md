# twitch-cli

A simple CLI Twitch tool, to use as an interactive middleware with streamlink, or youtube-dl, or anything else. It only spits out a string format of the Twitch stream url.

Example output:

`https://www.twitch.tv/rwxrob 160p`

Or for vods:

`https://www.twitch.tv/videos/757370887`

streamlink, or youtube-dl can then use this URL to run their commands. You can use command substitution, or whatever scripts you want to make those run it.

## Example usage

Currently there aren't that many flags.

Running just the binary requires to have the authentication set up. It shows you the

`-v` Puts you inside vod mode. Instead of listing streamers that are live, that you follow, it lists ALL streamers that you follow, even offline ones.

`-q` Decides if you want to pick qualities or not

## TODO

 [x] List information and pick following live streams

 [] Format output to make important information stick out a bit more

 [] Support resetting of oauth key

 [x] Support picking of quality

 [x] Support VODs

 [] Full CLI support with flags

 [] Other Theatron in bash features

 [] Optional: support images using SIXEL?

 [] Don't rely on ktr0731's go fuzzy finder completely

 [] Refractor

 [] Combo mode, meaning we get to see all streamers we follow, their live status, then we get to pick if we want to watch their vods or the livestream

## Status

The project is in early development stages, and is currently only really useful to get information on currently live streams. I don't know Golang. I suck at it in fact, all constructive criticism is more than welcome in issues. Contributions of any sort are also very encouraged.

If you have any feature requests, don't hesitate to ask. If it's possible I'll hope to do it.

## Installation

Currently I don't recommend installing the software yet. But if you decide to help me develop, it relies on the following environment variables:

* THEATRON_OAUTH_KEY, which you can either export yourself, or put in your .env file. This is your oauth key for the Twitch account, so don't share it with anyone.

### On NixOS

Put this in your `environment.systemPackages` or `home.packages`:
```nix
(callPackage (fetchFromGitHub {
  owner = "davidv171";
  repo = "twitch-cli-menu";
  rev = "replace-this-with-commit-hash";
  sha256 = "0000000000000000000000000000000000000000000000000000"; # replace this with actual sha256 on fail
}) {})
```

## Early demo

[![DEMO](Early demo)](theatron-go-demo.webm)


