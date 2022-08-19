import { ApplicationCommandType, ContextMenuCommandBuilder } from 'discord.js'
import { ContextMenu } from '../../structures/Interaction'

export default new ContextMenu(
  'sayMessage',
  new ContextMenuCommandBuilder()
    .setName('sayMessage')
    .setNameLocalization('ko', '이 메세지 읽기')
    .setType(ApplicationCommandType.Message)
    .toJSON(),
  async (client, interaction) => {
    if (!interaction.isMessageContextMenuCommand()) return
    interaction.reply(interaction.targetMessage.content)
  }
)
