package controller

import (
    "github.com/ishiikurisu/logey"
    "github.com/ishiikurisu/logeybot/model"
    "github.com/ishiikurisu/logeybot/view"
    "strings"
    "strconv"
    "fmt"
)

type Controller struct {
    ID int64
    Logey logey.Log
    View view.Conversation
}

// Creates a new controller for that id.
func NewController(inlet int64) Controller {
    // TODO Load previous Logey from memory if this ID already has a database entry
    storedLog, _ := model.LoadLog(inlet)
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
    } else if strings.HasPrefix(message, "/add") || strings.HasPrefix(message, "/start") {
        controller.View = view.NewAdditionConversation()
        outlet = controller.View.Speak()
    } else if strings.HasPrefix(message, "/get") {
        outlet = view.Prettify(controller.Logey.ToString())
    } else if strings.HasPrefix(message, "/money") {
        outlet = fmt.Sprintf("%.2f$", controller.Logey.CalculateBalance())
    } else if strings.HasPrefix(message, "/export") {
        outlet = model.GetIdFile(controller.ID)
    } else {
        outlet = "wtf?"
    }

    return outlet
}

func (controller *Controller) BeUp(message string) string {
    outlet := ""

    controller.View.Listen(message)
    if !controller.View.IsUp() {
        outlet = controller.Dump()
        controller.View = view.CreateEmptyConversation()
    } else {
        outlet = controller.View.Speak()
    }

    return outlet
}

// Creates a new entry and saves it to memory
func (controller *Controller) Dump() string {
    message := ""
    answers := controller.View.Answers
    what := answers[0]
    howMuch, oops := strconv.ParseFloat(answers[1], 64)

    if oops != nil {
        message = "Invalid input!"
    } else {
        controller.Logey.Add(what, howMuch)
        model.SaveLog(controller.ID, controller.Logey.ToString())
        message = "Data saved on memory!"
    }

    return message
}

/* RANDOM FUNCS */

// Checks if the message to be sent is text or file
func GetMessageKind(message string) string {
    kind := "text"
    if strings.HasPrefix(message, "/export") {
        kind = "file"
    }
    return kind
}
