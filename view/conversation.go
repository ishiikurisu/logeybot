package view

// Defines a conversation to be performed between the user and the bot.
type Conversation struct {
    Steps []string
    Answers []string
    Topic int
    Kind string
}

/* ######################
   # EMPTY CONVERSATION #
   ###################### */

// Creates an empty view
func CreateEmptyConversation() Conversation {
    talk := Conversation {
        Steps: nil,
        Answers: nil,
        Topic: -1,
    }
    return talk
}

// Checks if a conversation is empty
func (talk *Conversation) IsEmpty() bool {
    return talk.Steps == nil
}

/* ######################
   # CONVERSATION LOGIC #
   ###################### */

// Creates a new conversation
func NewConversation(steps []string) Conversation {
    talk := Conversation {
        Steps: steps,
        Answers: make([]string, len(steps)),
        Topic: 0,
    }
    return talk
}

// Check if there are any more steps to this conversation
func (talk *Conversation) IsUp() bool {
    return (len(talk.Steps) > talk.Topic) && (!talk.IsEmpty())
}

// Gives the next question
func (talk *Conversation) Speak() string {
    return talk.Steps[talk.Topic]
}

// Listen to the answer to the question and moves the next one
func (talk *Conversation) Listen(answer string) string {
    talk.Answers[talk.Topic] = answer
    talk.Topic++
    return answer
}

/* ##########################
   # SPECIFIC CONVERSATIONS #
   ########################## */

// Creates an addition conversation
func NewAdditionConversation() Conversation {
    questions := []string {
        "What?",
        "How much?",
    }
    return NewConversation(questions)
}
