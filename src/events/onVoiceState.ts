import { ChannelType } from 'discord.js'
import { Event } from '../structures/Event'
import Embed from '../utils/Embed'

export default new Event(
  'voiceStateUpdate',
  async (client, oldState, newState) => {
    const embed = new Embed(client, 'info')
    if (oldState.channelId === newState.channelId) return

    if (oldState && !oldState.member?.user.bot) {
      embed
        .setAuthor({
          name: oldState.member?.user.tag ?? '알수없음',
          iconURL: oldState.member?.user.displayAvatarURL()
        })
        .setTitle('소환사 한명이 음성채팅방에서 떠났습니다')

      oldState.channel?.type !== ChannelType.GuildStageVoice
        ? oldState.channel
            ?.send({ embeds: [embed] })
            .then((msg) => msg.react('👋'))
        : null

      if (oldState.channel?.members.size ?? 0 <= 1) {
        await oldState.guild.members.me?.voice.disconnect()
      }
    }

    if (newState && !newState.member?.user.bot) {
      embed
        .setAuthor({
          name: newState.member?.user.tag ?? '알수없음',
          iconURL: newState.member?.user.displayAvatarURL()
        })
        .setTitle(
          `야생에 ${
            newState.member?.nickname ?? newState.member?.user.username
          }이(가) 나타났다!`
        )

      newState.channel?.type !== ChannelType.GuildStageVoice
        ? newState.channel
            ?.send({ embeds: [embed] })
            .then((msg) => msg.react('👋'))
        : null
    }
  }
)
