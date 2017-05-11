class Entry
    attr_reader :what
    attr_reader :how_much

    def initialize args
        @what = args[:what]
        @how_much = args[:how_much]
    end
end
