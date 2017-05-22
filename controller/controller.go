package controller

import (
    "github.com/ishiikurisu/logey"
    "github.com/ishiikurisu/logeybot/model"
    "github.com/ishiikurisu/logeybot/view"
    "strings"
)

type Controller struct {
    ID int64
    Logey logey.Log
    // TODO Add a view entity here
    View view.Conversation
}

// Creates a new controller for that id.
func NewController(inlet int64) Controller {
    // TODO Load previous Logey from memory if this ID already has a database entry
    storedLog := model.LoadLog(inlet)
    c := Controller {
        ID: inlet,
        Logey: logey.LogFromString(storedLog),
        View: view.CreateEmptyConversation(),
    }
    return c
}

// Generates the correct answer depending on the current controller's state
// the command given by the user.
func (c *Controller) Listen(message string) string {
    outlet := ""

    if c.View.IsUp() {
        c.View.Listen(message)
        if !c.View.IsUp() {
            // TODO Save data on memory
            // QUESTION Will this ending depend on the conversation kind?
            outlet = "Data saved on memory!"
            c.View = view.CreateEmptyConversation()
        } else {
            outlet = c.View.Speak()
        }
    } else if strings.HasPrefix(message, "/add") {
        c.View = view.NewAdditionConversation()
        outlet = c.View.Speak()
    } else {
        outlet = "wtf?"
    }

    return outlet
}
