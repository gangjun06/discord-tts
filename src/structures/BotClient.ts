import { Client, ClientOptions, ClientEvents, Collection } from 'discord.js'
import Logger from '../utils/Logger'

import { BaseCommand, Event } from '../../types/structures'
import config from '../config'
import CommandManager from '../managers/CommandManager'
import EventManager from '../managers/EventManager'
import ErrorManager from '../managers/ErrorManager'
import { config as dotenvConfig } from 'dotenv'
import { BaseInteraction } from './Interaction'
import Dokdo from 'dokdo'
import InteractionManager from '../managers/InteractionManager'
import * as Utils from '../utils/Utils'

const logger = new Logger('bot')

export default class BotClient extends Client {
  public readonly VERSION: string
  public readonly BUILD_NUMBER: string
  public readonly config = config

  public commands: Collection<string, BaseCommand> = new Collection()
  public events: Collection<keyof ClientEvents, Event> = new Collection()
  public errors: Collection<string, string> = new Collection()
  public interactions: Collection<string, BaseInteraction> = new Collection()
  public utils = Utils
  public command: CommandManager = new CommandManager(this)
  public event: EventManager = new EventManager(this)
  public error: ErrorManager = new ErrorManager(this)
  public interaction: InteractionManager = new InteractionManager(this)
  public dokdo: Dokdo = new Dokdo(this, {
    prefix: this.config.bot.prefix,
    noPerm: async (message) =>
      message.reply('You do not have permission to use this command.')
  })
  public constructor(options: ClientOptions) {
    super(options)
    dotenvConfig()

    logger.info('Loading config data...')

    this.VERSION = config.BUILD_VERSION
    this.BUILD_NUMBER = config.BUILD_NUMBER
  }

  public async start(token: string = config.bot.token): Promise<void> {
    logger.info('Logging in bot...')
    await this.login(token).then(() => this.setStatus())
  }

  public async setStatus(
    status: 'dev' | 'online' = 'online',
    name = '점검중...'
  ) {
    if (status.includes('dev')) {
      logger.warn('Changed status to Developent mode')
      this.user?.setPresence({
        activities: [
          { name: `${this.config.bot.prefix}도움 | ${this.VERSION} : ${name}` }
        ],
        status: 'dnd'
      })
    } else if (status.includes('online')) {
      logger.info('Changed status to Online mode')

      this.user?.setPresence({
        activities: [
          {
            name: `${this.config.bot.prefix}도움 | v${this.VERSION}(${this.BUILD_NUMBER})`
          }
        ],
        status: 'online'
      })
    }
  }
}
