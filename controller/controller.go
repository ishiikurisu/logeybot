package controller

type Controller struct {
    ID int64
    // TODO Add a Logey here
}

// Creates a new controller for that id.
func NewController(inlet int64) Controller {
    // TODO Load previous Logey from memory if this ID already has a database entry
    c := Controller {
        ID: inlet,
    }
    return c
}

// Generates the correct answer depending on the current controller's state
// the command given by the user.
func (c *Controller) Listen(message string) string {
    return message
}
