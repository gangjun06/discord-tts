package helper

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/discord-tts/bot/voice"
	"github.com/gangjun06/discord-tts/log"
	"github.com/gangjun06/discord-tts/tts"
)

var voiceHelpers map[string]*VoiceHelper

type Queue struct {
	Content string
	TtsType string
	Etc     string
}

type VoiceHelper struct {
	Voice    *voice.Voice
	Queue    []*Queue
	StopChan chan bool
}

func init() {
	voiceHelpers = make(map[string]*VoiceHelper)
}

func FindVoice(channelId string) *VoiceHelper {
	return voiceHelpers[channelId]
}

func SetVoice(channelId string, voiceConn *discordgo.VoiceConnection) *VoiceHelper {
	voice := voice.NewVoice()
	voice.SetVoiceConnection(voiceConn)
	helper := &VoiceHelper{
		Voice:    voice,
		Queue:    []*Queue{},
		StopChan: nil,
	}
	voiceHelpers[channelId] = helper
	return helper
}

func (v *VoiceHelper) AppendQueue(content, ttsType, etc string) {
	queue := &Queue{
		Content: content,
		TtsType: ttsType,
		Etc:     etc,
	}
	if len(v.Queue) == 0 && v.StopChan == nil {
		go v.play(queue)
		return
	}
	v.Queue = append(v.Queue, queue)
}

func (v *VoiceHelper) play(queue *Queue) {
	v.StopChan = make(chan bool)
	id := v.Voice.Vc.ChannelID
	if queue.TtsType == "kakao" {
		if err := tts.DownloadKakaoTTS(tts.VoiceType(queue.Etc), queue.Content, id); err != nil {
			log.Error(err)
			return
		}
	} else {
		if err := tts.SansVoice(queue.Content, id); err != nil {
			log.Error(err)
			return
		}
	}

	v.Voice.PlayAudioFile(tts.GetTempFilePath(id), v.StopChan)
	v.StopChan = nil
	v.onEnd()
}

func (v *VoiceHelper) onEnd() {
	var x *Queue
	if len(v.Queue) > 0 {
		x, v.Queue = v.Queue[0], v.Queue[1:]
		go v.play(x)
	}
}

func (v *VoiceHelper) Leave(disconnect bool) {
	if disconnect {
		v.Voice.Vc.Disconnect()
	}
}

func DeleteVoice(channelId string) {
	delete(voiceHelpers, channelId)
}
