package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type todo struct {
	Text string
}

func New(text string) *todo {
	return &todo{
		text,
	}
}

func (n todo) PrintNotes() {
	fmt.Printf("Text: %s\n", n.Text)
}

func (n *todo) SaveNotes() {
	fileName := strings.ReplaceAll(n.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	os.WriteFile(fileName, json, 0644)
}
