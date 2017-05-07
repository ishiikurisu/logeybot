require 'telegram_bot'

bot = TelegramBot.new token: ARGV[0]
bot.get_updates do |message|
  puts "@#{message.from.username}: #{message.text}"
  command = message.get_command_for bot

  message.reply do |reply|
    case command
    when /greet/i
      reply.text = "Hello, #{message.from.first_name}!"
    else
      reply.text = "#{message.from.first_name}, I have no idea what #{command.inspect} means."
    end
    puts "sending #{reply.text.inspect} to @#{message.from.username}"
    reply.send_with bot
  end
end
