import axios from 'axios'
import fetch from 'node-fetch'
import config from '../config'
import { KakaoVoiceType } from './Constants'

export const sendVoice = async (type: KakaoVoiceType, content: string) => {
  return await fetch(`https://kakaoi-newtone-openapi.kakao.com/v1/synthesize`, {
    method: 'POST',
    headers: {
      Authorization: `KakaoAK ${config.kakaoApiKey}`,
      'Content-Type': 'application/xml'
    },
    body: `<speak><voice name="${type}">${content}</voice></speak>`
  })
}
