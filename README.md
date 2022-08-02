# Discord TTS봇

Discord 음성채팅을 원활하게 도와주는 TTS봇입니다.
카카오의 TTS API를 이용하여 작동합니다.

## 사용법

```bash
git clone https://github.com/gangjun06/discord-tts
cd discord-tts
cp config.example.json config.json
# EDIT CONFIG
```

### Go

```bash
go build -o main .
./main
```

### Docker

```bash
docker build -t discord-tts .
docker run -v /path/to/config.json:/config.json discord-tts
```
