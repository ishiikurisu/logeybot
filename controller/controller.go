package controller

type Controller struct {
    ID int64
}

func NewController(inlet int64) Controller {
    c := Controller {
        ID: inlet,
    }
    return c
}

func (c *Controller) Listen(message string) string {
    return message
}
