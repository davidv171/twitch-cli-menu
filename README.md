# twitch-cli

A simple CLI Twitch tool, to use as an interactive middleware with streamlink, or youtube-dl, or anything else. It only spits out a string format of the Twitch stream url.


## TODO

[x] List information and pick following live streams
[] Format output to make important information stick out a bit more
[] Support resetting of oauth key
[] Support picking of quality
[] Support VODs
[] Full CLI support with flags
[] Other Theatron in bash features
[] Optional: support images using SIXEL?
[] Don't rely on ktr0731's go fuzzy finder completely
[] Refractor

## Status

The project is in early development stages, and is currently only really useful to get information on currently live streams. I don't know Golang. I suck at it in fact, all constructive criticism is more than welcome in issues. Contributions of any sort are also very encouraged.

If you have any feature requests, don't hesitate to ask. If it's possible I'll hope to do it.

## Installation

Currently I don't recommend installing the software yet. But if you decide to help me develop, it relies on the following environment variables:

* THEATRON_OAUTH_KEY, which you can either export yourself, or put in your .env file. This is your oauth key for the Twitch account, so don't share it with anyone.


## Early demo

[![DEMO](Early demo)](theatron-go-demo.webm)


