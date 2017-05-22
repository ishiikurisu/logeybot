package view

import "testing"

func TestCanCreateANewConversation(t *testing.T) {
    questions := []string {
        "how are you?",
        "how old are you?",
    }
    talk := NewConversation(questions)
    if len(talk.Steps) != len(talk.Answers) {
        t.Fatalf("Invalid talk structure")
    }
}

func TestConversationFlow(t *testing.T) {
    questions := []string {
        "how are you?",
        "how old are you?",
    }
    talk := NewConversation(questions)
    for talk.IsUp() {
        talk.Speak()
        talk.Listen("meh")
    }
    for _, answer := range talk.Answers {
        if answer != "meh" {
            t.Fatalf("Not all questions answered")
        }
    }
}

func TestIfEmptyConversationKeepsStopped(t *testing.T) {
    blah := CreateEmptyConversation()
    if blah.IsUp() {
        t.Fatalf("Empty talk is up #wtf")
    }
}
