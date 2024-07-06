package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "GET", Description: ""},
		{Text: "POST", Description: ""},
		{Text: "PUT", Description: ""},
	}
	currentText := d.Text
	fmt.Printf("current: %s\n", currentText)

	if strings.HasPrefix(currentText, "GET ") {
		s = []prompt.Suggest{
			{Text: "/", Description: ""},
			{Text: "/aaa", Description: ""},
			{Text: "/bbb", Description: ""},
		}
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}
