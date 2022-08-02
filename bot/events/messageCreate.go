package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/bot/commands"
	"github.com/gangjun06/discord-tts/config"
	"github.com/gangjun06/discord-tts/tts"
)

func MessageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}

	if !strings.HasPrefix(msg.Content, config.Get().Prefix) {
		return
	}

	splited := strings.Split(msg.Content, " ")
	command := strings.TrimPrefix(splited[0], config.Get().Prefix)

	voiceType := tts.VoiceManBright

	switch command {
	case "2":
		voiceType = tts.VoiceWomanBright
	case "3":
		voiceType = tts.VoiceManCalm
	case "4":
		voiceType = tts.VoiceWomanCalm
	case "0":
		voiceType = "sans"
	case "멈춰":
		// if stopChan != nil {
		// 	stopChan <- true
		// }
		return
	// case "제비뽑기":
	case "나가":
		commands.Leave(s, msg)
		return
	case "도움":
		commands.DisplayHelp(s, msg.ChannelID)
		return
	}

	if len(splited) < 1 {
		return
	}

	commands.TTS(s, msg, splited[1:], voiceType)
}
