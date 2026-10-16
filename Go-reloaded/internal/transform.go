package internal

import (
	"errors"
	"strconv"
	"strings"
)

func Hex(text []rune) ([]rune, error) {
	number := 0
	for _, character := range text {
		character := strings.ToUpper(string(character))
		if !strings.Contains("0123456789ABCDEF", character) {
			return []rune{}, errors.New("L'élément rentrer en paramètres n'est pas un nombre hexadécimale")
		}
		number = number*16 + strings.Index("0123456789ABCDEF", character)
	}
	return []rune(strconv.Itoa(number)), nil
}

func Bin(text []rune) ([]rune, error) {
	number := 0
	for _, digit := range text {
		if digit == '0' {
			number *= 2
		} else if digit == '1' {
			number = number*2 + 1
		} else {
			return []rune{}, errors.New("L'élément rentrer en paramètres n'est pas un nombre binaire")
		}
	}
	return []rune(strconv.Itoa(number)), nil
}

func Up(text []rune) []rune {
	for index, character := range text {
		text[index] = []rune(strings.ToUpper(string(character)))[0]
	}
	return text
}

func Low(text []rune) []rune {
	for index, character := range text {
		text[index] = []rune(strings.ToLower(string(character)))[0]
	}
	return text
}

func Cap(text []rune) []rune {
	text[0] = []rune(strings.ToUpper(string(text[0])))[0]
	return text
}
