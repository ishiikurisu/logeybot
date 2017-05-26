package controller

import (
    "github.com/ishiikurisu/logey"
    "github.com/ishiikurisu/logeybot/model"
    "github.com/ishiikurisu/logeybot/view"
    "strings"
    "strconv"
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
func (controller *Controller) Listen(message string) string {
    outlet := ""

    if controller.View.IsUp() {
        if strings.HasPrefix(message, "/cancel") {
            controller.View = view.CreateEmptyConversation()
            outlet = "Operation cancelled!"
        } else {
            outlet = controller.BeUp(message)
        }
    } else if strings.HasPrefix(message, "/add") {
        controller.View = view.NewAdditionConversation()
        outlet = controller.View.Speak()
    } else if strings.HasPrefix(message, "/get") {
        // TODO Add retrieve log logic
        // TODO Prettify log before displaying it
        outlet = controller.Logey.ToString()
    } else {
        outlet = "wtf?"
    }

    return outlet
}

func (controller *Controller) BeUp(message string) string {
    outlet := ""

    controller.View.Listen(message)
    if !controller.View.IsUp() {
        controller.Dump()
        // QUESTION Will this ending depend on the conversation kind?
        outlet = "Data saved on memory!"
        controller.View = view.CreateEmptyConversation()
    } else {
        outlet = controller.View.Speak()
    }

    return outlet
}

// Creates a new entry and saves it to memory
func (controller *Controller) Dump() {
    answers := controller.View.Answers
    what := answers[0]
    howMuch, oops := strconv.ParseFloat(answers[1], 64)
    if oops != nil {
        panic(oops)
    }
    controller.Logey.Add(what, howMuch)
    // TODO Save logey on memory
}
