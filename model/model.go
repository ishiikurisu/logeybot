package model

import (
	"fmt"
	"github.com/ishiikurisu/logey"
	"io/ioutil"
	"os"
)

type Model struct {
	Id  int64
	Log *logey.Log
}

// Loads an account from memory
func NewModel(id int64) *Model {
	rawLog, oops := loadLog(id)
	log := logey.NewLog()
	if oops == nil {
		log = logey.Import(rawLog)
	}
	m := Model{
		Id:  id,
		Log: log,
	}
	return &m
}

// Updates account with new message
func (model *Model) Update(message string) error {
	entry, oops := logey.Understand(message)
	if oops != nil {
		return oops
	}
	model.Log.AddEntry(entry)
	oops = saveLog(model.Id, model.Log.Export())
	if oops != nil {
		// XXX what if storing stuff fails?
		return oops
	}
	return nil
}

// Gets the account balance
func (model Model) GetBalance() float64 {
	return model.Log.Balance
}

// Gets the entries
func (model Model) GetDescriptions() []string {
	// TODO Refactor this function
	outlet := make([]string, 0)
	for _, entry := range model.Log.Entries {
		outlet = append(outlet, fmt.Sprintf("%s: %.2f", entry.How, entry.HowMuch))
	}
	return outlet
}

// Gets the full file name, including the data directory
func getIdFile(id int64) string {
	return fmt.Sprintf("./data/logeybot/%d.jsonl", id)
}

// Loads a raw log from memory based on the user id.
func loadLog(id int64) (string, error) {
	raw := ""
	target := getIdFile(id)

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

// Saves log in memory
func saveLog(id int64, log string) error {
	target := getIdFile(id)
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
