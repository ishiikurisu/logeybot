package model

import (
    "fmt"
    "io/ioutil"
)

// Loads a raw log from memory based on the user id.
func LoadLog(id int64) string {
    raw := ""
    files, _ := ioutil.ReadDir("data/logeybot")
    target := BuildTargetName(id)

    for _, file := range files {
        if file.Name() == target {
            content, _ := ioutil.ReadFile(target)
            raw = string(content)
        }
    }

    return raw
}

// Generates the file name for a respective user id.
func BuildTargetName(id int64) string {
    return fmt.Sprintf("%d.txt", id)
}

// TODO Implement procedure to save log
