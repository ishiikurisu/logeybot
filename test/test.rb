require "minitest/autorun"

class TestPerson < MiniTest::Test
    def setup
        # this is run before every test
    end

    def test_get_files
        # assert_equal 5, @jukebox.get_available_songs.length
    end

    def test_first_song
        # assert_empty @jukebox.play_nth_song 0
    end
end
