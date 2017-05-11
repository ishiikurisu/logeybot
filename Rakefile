task default: %w[run]

desc "Starts the bot"
task :run do |t, args|
    ruby "app/main.rb #{ENV['API']}"
end

desc "Unit tests"
task :test do
    ruby 'test/test_entry.rb'
end
