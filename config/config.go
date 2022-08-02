package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gangjun06/discord-tts/log"
)

type ConfigBot struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

type Config struct {
	KakaoApiKey    string      `json:"kakaoApiKey"`
	Prefix         string      `json:"prefix"`
	Bots           []ConfigBot `json:"bots"`
	VoiceJoinAlert bool        `json:"voiceJoinAlert"`
}

type Effects struct {
	BaseURL string              `json:"baseURL"`
	Data    map[string][]string `json:"data"`
}

var (
	config *Config
)

func init() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Failed to read Config.json file")
	}
	if err := json.Unmarshal(file, &config); err != nil {
		log.Fatal("Failed to parse config.json file")
	}
}

func Get() *Config {
	return config
}
