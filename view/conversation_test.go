package view

import "testing"

func TestCanCreateANewConversation(t *testing.T) {
    questions := []string {
        "how are you?",
        "how old are you?",
    }
    talk := NewConversation(questions)
    if len(talk.Steps) != len(talk.Answers) {
        t.Fatalf("TODO Give a more meaningful log here")
    }
}

// TODO Make bot listen to stuff and move from topic to topic
