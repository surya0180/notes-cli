package notes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type notes struct {
	Title   string
	Content string
}

func New(title string, content string) *notes {
	return &notes{
		title,
		content,
	}
}

func (n notes) PrintNotes() {
	fmt.Printf("Title: %s\n", n.Title)
	fmt.Printf("Content: \n%s\n", n.Content)
}

func (n *notes) SaveNotes() {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	os.WriteFile(fileName, json, 0644)
}
