import { SlashCommandBuilder } from 'discord.js'
import { BaseCommand } from '../../structures/Command'
import { helpContent } from '../../utils/Constants'
import Embed from '../../utils/Embed'

export default new BaseCommand(
  {
    name: 'help',
    description: '지금 이 도움말을 출력합니다',
    aliases: ['ehdna', 'ehdnaakf', '도움말', '도움']
  },
  async (client, message) => {
    const embed = new Embed(client, 'info')
      .setTitle('Discord TTS - 기본 도움말')
      .setDescription(helpContent)

    message.reply({
      embeds: [embed]
    })
  },
  {
    data: new SlashCommandBuilder()
      .setName('help')
      .setNameLocalization('ko', '도움말')
      .setDescription('지금 이 도움말을 출력합니다')
      .toJSON(),
    async execute(client, interaction) {
      const embed = new Embed(client, 'info')
        .setTitle('Discord TTS - 기본 도움말')
        .setDescription(helpContent)

      interaction.reply({
        embeds: [embed],
        ephemeral: true
      })
    }
  }
)
