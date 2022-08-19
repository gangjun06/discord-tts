import 'dotenv/config'
import path from 'path'
import Logger from './utils/Logger'
import config from './config'

import BotClient from './structures/BotClient'
import CommandManager from './managers/CommandManager'
import EventManager from './managers/EventManager'
import InteractionManager from './managers/InteractionManager'

const logger = new Logger('main')

logger.log('Starting up...')

process.on('uncaughtException', (e) => logger.error(e.stack as string))
process.on('unhandledRejection', (e: Error) => logger.error(e.stack as string))

const client = new BotClient(config.bot.options)
const command = new CommandManager(client)
const event = new EventManager(client)
const interaction = new InteractionManager(client)

command.load(path.join(__dirname, 'commands'))
event.load(path.join(__dirname, 'events'))
interaction.load(path.join(__dirname, 'interactions'))

client.start(config.bot.token)
