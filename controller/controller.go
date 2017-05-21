package controller

import (
    "github.com/ishiikurisu/logey"
    "github.com/ishiikurisu/logeybot/model"
    "strings"
)

type Controller struct {
    ID int64
    Logey logey.Log
    // TODO Add a view entity here
}

// Creates a new controller for that id.
func NewController(inlet int64) Controller {
    // TODO Load previous Logey from memory if this ID already has a database entry
    storedLog := model.LoadLog(inlet)
    c := Controller {
        ID: inlet,
        Logey: logey.LogFromString(storedLog),
    }
    return c
}

// Generates the correct answer depending on the current controller's state
// the command given by the user.
func (c *Controller) Listen(message string) string {
    outlet := ""

    if strings.HasPrefix(message, "/add") {
        // TODO Start add conversation
        outlet = "how about we add something?"
    } else {
        outlet = "wtf?"
    }

    return outlet
}
