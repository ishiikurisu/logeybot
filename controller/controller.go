package controller

import (
	"github.com/ishiikurisu/logeybot/model"
	"github.com/ishiikurisu/logeybot/view"
	"strings"
)

type PossibleState int

const (
	IDLE PossibleState = iota
	WAITING_DESCRIPTION
)

// This is the basic struct to
type Controller struct {
	Model *model.Model
	State PossibleState
}

// Creates a new controller for that id.
func NewController(id int64) Controller {
	c := Controller{
		Model: model.NewModel(id),
		State: IDLE,
	}
	return c
}

// Generates the correct answer depending on the current controller's state
// the command given by the user.
func (controller *Controller) Listen(message string) string {
	outlet := "sorry, I didn't understand that."

	switch controller.State {
	case IDLE:
		if strings.HasPrefix(message, "/cancel") {
			outlet = "I wasn't doing anything anyways..."
		} else if strings.HasPrefix(message, "/add") ||
			strings.HasPrefix(message, "/start") {
			controller.State = WAITING_DESCRIPTION
			outlet = "Tell me more about it..."
		} else if strings.HasPrefix(message, "/money") {
			outlet = view.BalanceMessage(controller.Model.GetBalance())
		} else if strings.HasPrefix(message, "/get") {
			outlet = view.LogMessage(controller.Model.GetDescriptions())
		}
		break

	case WAITING_DESCRIPTION:
		if strings.HasPrefix(message, "/cancel") {
			outlet = "Operation cancelled"
		} else {
			controller.State = IDLE
			if oops := controller.Model.Update(message); oops == nil {
				outlet = view.BalanceMessage(controller.Model.GetBalance())
			}
		}
		break
	}

	return outlet
}

// Checks if the message to be sent is text or file
func GetMessageKind(message string) string {
	kind := "text"
	if strings.HasPrefix(message, "/export") {
		kind = "file"
	}
	return kind
}
