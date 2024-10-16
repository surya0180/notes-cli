package main

import (
	"github.com/notescli/client"
	"github.com/notescli/notes"
)

func main() {
	title := client.Input("Notes title: ")
	content := client.Input("Notes content: ")

	note1 := notes.New(title, content)
	note1.PrintNotes()
	note1.SaveNotes()
}
