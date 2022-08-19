import { PrismaClientOptions } from '@prisma/client/runtime'
import {
  ButtonInteraction,
  ChatInputCommandInteraction,
  ClientOptions,
  ColorResolvable,
  ContextMenuCommandInteraction,
  HexColorString,
  Interaction,
  Message,
  MessageContextMenuCommandInteraction,
  ModalSubmitInteraction,
  SelectMenuInteraction,
  ShardingManagerOptions
} from 'discord.js'
import { ReportType } from '../src/utils/Constants'
import ObjectConfig from '../src/config/example.config.json'

export interface ErrorReportOptions {
  executer:
    | Message<true>
    | ChatInputCommandInteraction<'cached'>
    | ContextMenuCommandInteraction<'cached'>
    | SelectMenuInteraction<'cached'>
    | ButtonInteraction<'cached'>
    | ModalSubmitInteraction<'cached'>
    | undefined
  isSend?: boolean
}

export type Config = {
  BUILD_VERSION: string
  BUILD_NUMBER: string
  devGuildID: string
  githubToken?: string
  name: string
  repository?: string
} & { logger: LoggerConfig } & { bot: BotConfig } & {
  report: ErrorReportConfig
} & Omit<typeof ObjectConfig, 'bots' | 'prefix'>

export interface LoggerConfig {
  level: LevelType
  dev: boolean
}

export interface ErrorReportConfig {
  type: ReportType
  webhook: {
    url: string
  }
  text: {
    guildID: string
    channelID: string
  }
}
export interface BotConfig {
  sharding: boolean
  shardingOptions?: ShardingManagerOptions
  options: ClientOptions
  token: string
  owners?: string[]
  prefix: string
}

export type DatabaseType = 'mongodb' | 'prisma'

export type LevelType =
  | 'fatal'
  | 'error'
  | 'warn'
  | 'info'
  | 'verbose'
  | 'debug'
  | 'chat'

export type EmbedType =
  | 'default'
  | 'error'
  | 'success'
  | 'warn'
  | 'info'
  | HexColorString

export * from './structures'
export * from './command'
