package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	if prompt == "" {
		panic("prompt is required")
	}

	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	if text == "" {
		panic("input cannot be empty")
	}

	return text
}
