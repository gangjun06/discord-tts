package tts

import (
	"os/exec"
	"strings"
)

func SansVoice(str, id string) error {
	list := "sox "
	for _, d := range strings.Split(str, "") {
		if d == " " {
			list += SansFilePath + "delay2.wav "
		} else {
			list += SansFilePath + "ay.wav " + SansFilePath + "delay.wav "
		}
	}
	list += SansFilePath + "delay.wav " + GetTempFilePath(id)

	cmd := exec.Command("/bin/sh", "-c", list)

	return cmd.Run()
}
