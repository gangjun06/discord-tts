package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/bot/helper"
	"github.com/gangjun06/discord-tts/config"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateGameStatus(0, config.Get().Prefix+"도움")
	helper.SetSession(s)
}
