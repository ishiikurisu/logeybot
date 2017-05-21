package view

// Defines a conversation to be performed between the user and the bot.
type Conversation struct {
    Steps []string
    Answers []string
    Topic int
}

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
    return len(talk.Steps) <= talk.Topic
}

// Gives the next question
func (talk *Conversation) Speak() string {
    return talk.Steps[0]
}

// Listen to the answer to the question and moves the next one
func (talk *Conversation) Listen(answer string) string {
    // TODO Implement listening algorithm
    return answer
}
