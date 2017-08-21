package model

import (
    "fmt"
    "io/ioutil"
    "os"
)

// Gets the data directory
func getDataDir() string {
    return "./data/logeybot"
}

// Gets the full file name, including the data directory
func GetIdFile(id int64) string {
    return fmt.Sprintf("%s/%d.txt", getDataDir(), id)
}

// Loads a raw log from memory based on the user id.
func LoadLog(id int64) (string, error) {
    raw := ""
    target := GetIdFile(id)

    if _, oops := os.Stat(target); oops != nil {
        return raw, oops
    }
    if content, oops := ioutil.ReadFile(target); oops == nil {
        raw = string(content)
    } else {
        return raw, oops
    }

    return raw, nil
}

// TODO Implement procedure to save log
func SaveLog(id int64, log string) error {
    target := GetIdFile(id)
    file, oops := os.Create(target)

    if oops != nil {
        return oops
    } else {
        defer file.Close()
    }

    fmt.Fprintf(file, "%s", log)
    file.Sync()

    return nil
}
