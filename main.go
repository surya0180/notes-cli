package main

import (
	"github.com/notescli/arrslice"
	"github.com/notescli/funks"
)

type saver interface {
	Save() error
}

func saveNotes(data saver) {
	data.Save()
}

func main() {
	arrslice.Execute()
	funks.Execute()
}
