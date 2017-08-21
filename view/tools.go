package view

import (
    "fmt"
    "bytes"
    "strings"
)

// Turns a raw log into a meaningful string
func Prettify(inlet string) string {
    var buffer bytes.Buffer

    for _, it := range strings.Split(inlet, "|") {
        if (len(it) > 0) && (it != "---") && (it != "...") {
            buffer.WriteString(fmt.Sprintf("%s\n", it))
        }
    }

    outlet := buffer.String()
    if len(outlet) == 0 {
        outlet = "Nothing added yet!"
    }
    return outlet
}
