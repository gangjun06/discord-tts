import { SlashCommandBuilder } from '@discordjs/builders'
import type {
  Message,
  ClientEvents,
  ChatInputCommandInteraction,
  Interaction,
  RESTPostAPIApplicationCommandsJSONBody
} from 'discord.js'
import BotClient from '../src/structures/BotClient'

export interface MessageCommnad {
  data: MessageCommandOptions
  execute: MessageCommandFuntion
}
export interface Command extends MessageCommnad {
  isSlash?: boolean
  slash?: SlashCommand
}

export interface SlashCommand {
  data: InteractionData
  execute: SlashCommandFunction
  options?: SlashCommandOptions
  slash?: SlashCommand
}

export interface MessageCommandOptions {
  name: string
  description?: string
  aliases: string[]
}

export type MessageCommandFuntion = (
  client: BotClient,
  message: Message<true>,
  args: string[]
) => Promise<any> | Promise<any>

export type SlashCommandFunction = (
  client: BotClient,
  interaction: ChatInputCommandInteraction
) => Promise<any>

export interface SlashCommandOptions {
  name: string
  isSlash?: boolean
}

export interface Event {
  name: keyof ClientEvents
  options?: EventOptions
  execute: (
    client: BotClient,
    ...args: ClientEvents[keyof ClientEvents]
  ) => Promise<any>
}

export type EventFunction<E extends keyof ClientEvents> = (
  client: BotClient,
  ...args: ClientEvents[E]
) => Promise<any>

export interface EventOptions {
  once: boolean
}

export type BaseInteractionFunction<T = Interaction> = (
  client: BotClient,
  interaction: T
) => Promise<any>

export type BaseCommand = MessageCommnad | SlashCommand | Command
export type InteractionData = RESTPostAPIApplicationCommandsJSONBody
