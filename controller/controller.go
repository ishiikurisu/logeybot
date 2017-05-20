package controller

import (
    "github.com/ishiikurisu/logey"
    "github.com/ishiikurisu/logeybot/model"
)

type Controller struct {
    ID int64
    // TODO Add a Logey here
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
    // TODO Implement a view entity to deal with listening to messages
    return message
}