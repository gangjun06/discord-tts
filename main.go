package main

import (
	"os"
	"os/signal"

	"github.com/gangjun06/discord-tts/bot"
)

func main() {
	bot.InitBot()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
