require "./lib/entry.rb"
require "minitest/autorun"

class TestEntry < MiniTest::Test
    def setup
        # this is run before every test
    end

    def test_entry_stores_data
        entry_name = "cake"
        entry_value = 2.00
        entry = Entry.new what: entry_name, how_much: entry_value
        assert_equal entry_name, entry.what
        assert_equal entry_value, entry.how_much
    end
end
