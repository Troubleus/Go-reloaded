package internal

import (
	"strings"
)

func TextToSlice(text string) [][][]rune {
	lines := strings.Split(text, "\n")
	newLines := make([][][]rune, 0)
	for _, line := range lines {
		line := strings.Split(line, " ")
		newLine := make([][]rune, 0)
		for _, word := range line {
			newLine = append(newLine, []rune(word))
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func SliceToText(line [][]rune) string {
	var text string
	for index, word := range line {
		if index == len(line)-1 {
			text += string(word)
		}else{
			text += string(word) + " "
		}
	}
	return text
}
