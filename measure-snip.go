package main

import (
	"encoding/json"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	type measurement struct {
		Revision  string `json:"revision"`
		Name      string `json:"name"`
		Input     string `json:"input"`
		Indicator string `json:"indicator"`
		Target    int    `json:"target"`
		Output    string `json:"output"`
		Width     int    `json:"width"`
	}
	var measurements []measurement
	for _, indicator := range []struct{ name, value string }{
		{"Plain", "..."},
		{"ANSI", "\x1b[33m...\x1b[0m"},
	} {
		for _, target := range []int{1, 2} {
			output := text.Snip("abcdef", target, indicator.value)
			measurements = append(measurements, measurement{
				Revision: os.Args[1], Name: indicator.name, Input: "abcdef",
				Indicator: indicator.value, Target: target, Output: output,
				Width: text.StringWidthWithoutEscSequences(output),
			})
		}
	}
	if err := json.NewEncoder(os.Stdout).Encode(measurements); err != nil {
		panic(err)
	}
}
