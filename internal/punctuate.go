package internal

import (
	"strings"
)

func FormatPunctuate(text [][]rune)[][]rune{
	textCopy:= [][]rune{}

	isApostrophe := 0
	for index,word := range text{
		if isApostrophe == 1 {
			isApostrophe = 2
			continue
		}
		if strings.Contains(".,!?:;",string(word[0])){
			textCopy[len(textCopy)-1] = append(textCopy[len(textCopy)-1],word[:]...)
		}else if word[0] == '\''{
			if isApostrophe == 0 {
				textCopy = append(textCopy, append([]rune{word[0]}, text[index+1][:]...))
				isApostrophe = 1
			}  else if isApostrophe == 2 {
				textCopy[len(textCopy)-1] = append(textCopy[len(textCopy)-1],word[:]...)
				isApostrophe = 0
			}
		} else{
			textCopy = append(textCopy, word)
		}
	}
	return textCopy
}