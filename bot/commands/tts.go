package commands

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/bot/helper"
	"github.com/gangjun06/discord-tts/tts"
)

func TTS(s *discordgo.Session, msg *discordgo.MessageCreate, content []string, voiceType tts.VoiceType) {

	reply := func(str string) {
		s.ChannelMessageSendReply(msg.ChannelID, str, msg.Reference())
	}

	userVoice, _ := s.State.VoiceState(msg.GuildID, msg.Author.ID)
	if userVoice == nil {
		reply("먼저 음성채팅방에 접속하여 주세요.")
		return
	}

	botVoice := helper.FindVoice(userVoice.ChannelID)

	if botVoice == nil || botVoice.Voice.Vc.ChannelID != userVoice.ChannelID {
		s2 := helper.NotBusyBot(msg.GuildID)
		if s2 == nil {
			reply("사용 가능한 봇이 존재하지 않습니다.")
			return
		}
		voiceConn, _ := s2.ChannelVoiceJoin(msg.GuildID, userVoice.ChannelID, false, true)

		botVoice = helper.SetVoice(userVoice.ChannelID, voiceConn)
	}

	joined := strings.Join(content, " ")

	if voiceType == "sans" {
		botVoice.AppendQueue(joined, "sans", "")
		return
	}

	if match, _ := regexp.MatchString("[가-힇]", joined); !match {
		return
	}

	botVoice.AppendQueue(joined, "kakao", string(voiceType))
}

func Leave(s *discordgo.Session, msg *discordgo.MessageCreate) {

	reply := func(str string) {
		s.ChannelMessageSendReply(msg.ChannelID, str, msg.Reference())
	}

	userVoice, _ := s.State.VoiceState(msg.GuildID, msg.Author.ID)
	if userVoice == nil {
		reply("먼저 음성채팅방에 접속하여 주세요.")
		return
	}

	botVoice := helper.FindVoice(userVoice.ChannelID)

	if botVoice == nil {
		reply("봇과 같은 음성채팅방에 접속하여 주세요.")
		return
	}
	botVoice.Voice.Vc.Disconnect()
	s.MessageReactionAdd(msg.ChannelID, msg.ID, "✅")
}
