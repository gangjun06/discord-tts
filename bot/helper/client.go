package helper

import "github.com/bwmarrin/discordgo"

var sessions map[string]*discordgo.Session

func init() {
	sessions = make(map[string]*discordgo.Session)
}

func SetSession(s *discordgo.Session) {
	sessions[s.State.User.ID] = s
}

func NotBusyBot(guildId string) *discordgo.Session {
	for _, d := range sessions {
		if d.VoiceConnections[guildId] == nil {
			return d
		}
	}
	return nil
}
