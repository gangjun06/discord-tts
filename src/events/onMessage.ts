import { Event } from '../structures/Event'
import CommandManager from '../managers/CommandManager'
import ErrorManager from '../managers/ErrorManager'
import type { MessageCommand } from '../structures/Command'
import { sendVoice } from '../utils/Utils'
import { KakaoVoiceType } from '../utils/Constants'
import {
  createAudioPlayer,
  createAudioResource,
  joinVoiceChannel,
  NoSubscriberBehavior,
  StreamType
} from '@discordjs/voice'
import fs from 'fs'
import { Readable } from 'stream'

export default new Event('messageCreate', async (client, message) => {
  const commandManager = new CommandManager(client)
  const errorManager = new ErrorManager(client)

  if (message.author.bot) return
  if (!message.inGuild()) return
  if (!message.content.startsWith(client.config.bot.prefix)) return

  const args = message.content
    .slice(client.config.bot.prefix.length)
    .trim()
    .split(/ +/g)
  const commandName = args.shift()?.toLowerCase()
  const command = commandManager.get(commandName as string) as MessageCommand

  await client.dokdo.run(message)

  if (!command && commandName !== 'dok' && commandName !== '애ㅏ') {
    const data = await sendVoice(KakaoVoiceType.ManBright, args.join(' '))

    const connection = joinVoiceChannel({
      channelId: message.member?.voice.channelId as string,
      guildId: message.guild.id,
      adapterCreator: message.channel.guild.voiceAdapterCreator
    })

    const player = createAudioPlayer({
      behaviors: {
        noSubscriber: NoSubscriberBehavior.Pause
      }
    })

    const res = createAudioResource(data.body, {
      type: StreamType.WebmOpus
    })

    player.play(res)
    connection.subscribe(player)
    client.connections.set(message.guild.id, connection)
  }
  try {
    await command?.execute(client, message, args)
  } catch (error: any) {
    errorManager.report(error, { executer: message, isSend: true })
  }
})
