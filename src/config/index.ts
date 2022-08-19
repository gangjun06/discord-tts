import { execSync } from 'child_process'
import { ReportType } from '../utils/Constants'
import { bots, kakaoApiKey, prefix, voiceJoinAlert } from './config.json'
import { Config } from '../../types'

const config: Config = {
  BUILD_NUMBER: execSync('git rev-parse --short HEAD').toString().trim(),
  BUILD_VERSION: '0.1.2',
  kakaoApiKey,
  voiceJoinAlert,
  devGuildID: '',
  githubToken: '',
  name: 'Discord TTS',
  bot: {
    sharding: false,
    shardingOptions: {},
    options: {
      intents: [130815],
      allowedMentions: { parse: ['users', 'roles'], repliedUser: false }
    },
    token: bots[0].token,
    owners: [],
    prefix
  },
  report: {
    type: ReportType.Webhook,
    webhook: {
      url: ''
    },
    text: {
      guildID: '',
      channelID: ''
    }
  },
  logger: {
    level: 'chat',
    dev: true
  }
}

export default config
