import { type Client, EmbedBuilder, type EmbedData } from 'discord.js'
import { EmbedType } from '../../types'

export default class Embed extends EmbedBuilder {
  constructor(client: Client, type: EmbedType) {
    if (!client.isReady()) return

    const EmbedJSON: EmbedData = {
      timestamp: new Date().toISOString()
    }

    super(EmbedJSON)

    if (type === 'success') this.setColor('#85FFBD')
    else if (type === 'error') this.setColor('#FF2525')
    else if (type === 'warn') this.setColor('#FBDA61')
    else if (type === 'info') this.setColor('#94bbe9')
    else if (type === 'default') this.setColor('#94bbe9')
    else this.setColor(type)
  }

  setType(type: EmbedType) {
    if (type === 'success') this.setColor('#85FFBD')
    else if (type === 'error') this.setColor('#FF2525')
    else if (type === 'warn') this.setColor('#FBDA61')
    else if (type === 'info') this.setColor('#94bbe9')
    else if (type === 'default') this.setColor('#94bbe9')
    else this.setColor(type)
  }
}
