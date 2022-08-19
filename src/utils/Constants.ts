import { inlineCode } from 'discord.js'

export const enum InteractionType {
  SlashCommand,
  Button,
  Select,
  ContextMenu,
  Modal,
  AutoComplete
}

export const enum ReportType {
  Webhook,
  Text
}

export const enum KakaoVoiceType {
  ManBright = 'MAN_DIALOG_BRIGHT',
  ManCalm = 'VOICE_MAN_CALM',
  WomanBright = 'WOMAN_DIALOG_BRIGHT',
  WomanCalm = 'WOMAN_READ_CALM'
}

export const helpContent = `
마이크없이 원할한 음성채팅을 도와줍니다.
> 봇 기본 사용법
${inlineCode(';도움')}: 지금 이 도움말을 출력합니다
${inlineCode(';! <하고싶은말>')}: 음성을 사용자 이름과 함께 재생합니다.
${inlineCode(';<목소리번호> <하고싶은말>')}: 음성채팅방에 TTS를 재생합니다.
${inlineCode(';나가')}: 봇을 음성채팅방에서 내보냅니다.

> 목소리번호 목록
1 - 남성 밝은 대화체
2 - 여성 밝은 대화체
3 - 남성 차분한 낭독체
4 - 여성 차분한 낭독체
` as const
