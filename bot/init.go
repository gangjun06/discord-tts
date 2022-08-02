package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/bot/events"
	"github.com/gangjun06/discord-tts/config"
	"github.com/gangjun06/discord-tts/log"
)

func initBot(isMain bool, bot config.ConfigBot) {
	botLog := log.WithBot(bot.Name)

	client, err := discordgo.New("Bot " + bot.Token)
	if err != nil {
		botLog.Fatal("Failed to create bot")
	}

	client.Identify.Intents = discordgo.IntentsGuildVoiceStates

	if isMain {
		client.Identify.Intents |= discordgo.IntentsGuildMessages
		client.AddHandler(events.MessageCreate)
	}

	client.AddHandler(events.VoiceStateUpdate(isMain))
	client.AddHandler(events.Ready)

	if err = client.Open(); err != nil {
		msg := "Failed to open bot connection"
		if isMain {
			botLog.Fatal(msg)
		}
		botLog.Error(msg)
	}
	botLog.Info("Bot Connected")
}

func InitBot() {
	bots := config.Get().Bots
	for i, bot := range bots {
		initBot(i == 0, bot)
	}
}
