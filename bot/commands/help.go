package commands

import (
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/config"
)

func DisplayHelp(s *discordgo.Session, channelID string) {
	prefix := config.Get().Prefix

	content := ""
	list := []string{}
	// for k := range config.Effects.Data {
	// 	list = append(list, k)
	// }

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	for _, d := range list {
		content += d + ", "
	}

	content = strings.TrimRight(content, ", ")

	s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "디스코드 TTS - 기본 도움말",
		Description: "마이크없이 원할한 음성채팅을 도와줍니다.",
		Color:       0xbedbe9,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "봇 기본 사용법",
				Value: prefix + "`도움`" + `지금 이 도움말을 출력합니다
					` + prefix + "`! <하고싶은말>`" + `: 음성을 사용자 이름과 함께 재생합니다.
					` + prefix + "`<목소리번호> <하고싶은말>`" + `: 음성채팅방에 TTS를 재생합니다.
					` + prefix + "`나가`" + `: 봇을 음성채팅방에서 내보냅니다.`,
			},
			{
				Name: "목소리번호 목록",
				Value: ` 1: 남성 밝은 대화체
					2: 여성 밝은 대화체
					3: 남성 차분한 낭독체
					4: 여성 차분한 낭독체
					0: 샌즈(효과음 사용 불가)`,
			},
		},
	})
	// s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
	// 	Title:       "디스코드 TTS - 효과음",
	// 	Description: "하고싶은말에 <효과음이름> 과 같이 입력하여 말하는도중 효과음을 추가할 수 있습니다.",
	// 	Color:       0xbedbe9,
	// 	Fields: []*discordgo.MessageEmbedField{
	// 		{
	// 			Name:  "효과음 목록",
	// 			Value: content,
	// 		},
	// 	},
	// })
}
