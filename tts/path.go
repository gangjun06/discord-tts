package tts

const (
	SansFilePath = "__files/sans/"
	TempFilePath = "__temp/"
)

func GetTempFilePath(id string) string {
	return TempFilePath + id + ".mp3"
}
