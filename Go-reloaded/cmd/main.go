package main

import (
	"fmt"
	"os"
	"go-reloaded/internal"
)

func main() {
	content, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := internal.TextToSlice(string(content))
	textSting := ""

	for _, line := range text{
		newLine,err := internal.FormatParameters(line)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		newLine = internal.FormatPunctuate(newLine)
		lineString := internal.SliceToText(newLine)
		textSting += lineString + "\n"
	}
	
	err = os.WriteFile(os.Args[2], []byte(textSting), 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
