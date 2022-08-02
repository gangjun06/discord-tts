package tts

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/gangjun06/discord-tts/config"
)

const (
	API_URL = "https://kakaoi-newtone-openapi.kakao.com/v1/synthesize"
)

var (
	effectRegex, _ = regexp.Compile(`^<[가-힇a-zA-Z]+>$`)
)

type VoiceType string

const (
	VoiceManBright   VoiceType = "MAN_DIALOG_BRIGHT"
	VoiceWomanBright VoiceType = "WOMAN_DIALOG_BRIGHT"
	VoiceManCalm     VoiceType = "VOICE_MAN_CALM"
	VoiceWomanCalm   VoiceType = "WOMAN_READ_CALM"
	VoiceSANS        VoiceType = "SANS"
)

// ConvertEffect convert effect to voice
// returns converted, not found effect
func ConvertEffect(text string) (string, string) {
	// splited := strings.Split(text, " ")
	content := ""
	// for _, d := range splited {
	// 	if matchString := effectRegex.MatchString(d); !matchString {
	// 		content += d + " "
	// 		continue
	// 	}
	// 	text := strings.TrimRight(strings.TrimLeft(d, "<"), ">")
	// 	effect, ok := config.Effects.Data[text]
	// 	if !ok {
	// 		return "", text
	// 	}

	// 	randomIndex := rand.Intn(len(effect))
	// 	pick := effect[randomIndex]

	// 	url := fmt.Sprintf(`<audio src="%s%s"/>`, config.Effects.BaseURL, pick)
	// 	content += url + " "
	// }
	return content, ""
}

func DownloadKakaoTTS(voiceType VoiceType, content, id string) error {
	text := fmt.Sprintf(`<speak><voice name="%s">%s</voice></speak>`, voiceType, content)

	reader := strings.NewReader(text)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, API_URL, reader)

	req.Header.Add("Authorization", "KakaoAK "+config.Get().KakaoApiKey)
	req.Header.Add("Content-Type", "application/xml")

	if err != nil {
		return fmt.Errorf("creating tts request")
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("sending tts request")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error read response")
	}
	if err := ioutil.WriteFile(GetTempFilePath(id), data, 06644); err != nil {
		return fmt.Errorf("write file")
	}

	return nil
}
