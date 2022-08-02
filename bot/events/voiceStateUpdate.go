package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/bot/helper"
	"github.com/gangjun06/discord-tts/config"
)

func VoiceStateUpdate(isMain bool) func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	return func(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
		if v.UserID == s.State.User.ID { // If Current Bot
			if v.BeforeUpdate == nil && v.ChannelID != "" { // Join
				return
			}

			voice := helper.FindVoice(v.BeforeUpdate.ChannelID)
			if voice == nil {
				return
			}

			if v.BeforeUpdate != nil && v.ChannelID != "" { // Move
				voice.Leave(true)
				helper.DeleteVoice(v.BeforeUpdate.ChannelID)
			} else if v.BeforeUpdate != nil && v.ChannelID == "" { // Leave
				helper.DeleteVoice(v.BeforeUpdate.ChannelID)
			}
			return
		}

		if isMain && config.Get().VoiceJoinAlert {
			name := v.Member.Nick
			if name == "" {
				name = v.Member.User.Username
			}
			embed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					IconURL: v.Member.AvatarURL("128"),
					Name:    name,
				},
			}

			send := func(channelId string, join bool) {
				if join {
					embed.Title = "맴버가 음성채팅방에 들어왔습니다."
					embed.Color = 0x57F287
				} else {
					embed.Title = "맴버가 음성채팅방에서 나갔습니다."
					embed.Color = 0xEB459E
				}
				s.ChannelMessageSendEmbed(channelId, embed)
			}

			if v.BeforeUpdate == nil && v.ChannelID != "" { // Join
				send(v.ChannelID, true)
			} else if v.BeforeUpdate != nil && v.ChannelID != "" { // Move
				send(v.BeforeUpdate.ChannelID, false)
				send(v.ChannelID, true)
			} else if v.BeforeUpdate != nil && v.ChannelID == "" { // Leave
				send(v.BeforeUpdate.ChannelID, false)
			}
		}

		if v.BeforeUpdate != nil && v.ChannelID == "" { // Leave
			guild, err := s.State.Guild(v.GuildID)
			if err != nil {
				return
			}
			count := -1
			for _, vs := range guild.VoiceStates {
				if vs.ChannelID == v.BeforeUpdate.ChannelID {
					count++
				}
			}
			if count < 1 {
				voice := helper.FindVoice(v.BeforeUpdate.ChannelID)
				if voice == nil {
					return
				}
				voice.Leave(true)
				helper.DeleteVoice(v.BeforeUpdate.ChannelID)
			}
		}

	}
}
