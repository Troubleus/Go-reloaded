package internal

import (
	"errors"
	"strings"
	"reflect"

)

func FormatParameters(text [][]rune) ([][]rune, error) {
	for index,word := range text{
		if index < len(text)-1{
			if (reflect.DeepEqual(word, []rune{'a'}) || reflect.DeepEqual(word, []rune{'A'})) && strings.Contains("aeiouhAEIOUH", string(text[index+1][0])) {
				text[index] = append(word, []rune{'n'}...)
			}
		}
		if word[0] == '('{
			parameter, number, err := Parameters(word)
			if err != nil {
				return [][]rune{}, err
			}
			switch parameter{
			case "hex":
				decimal,err := Hex(text[index-1])
				if err != nil {
					return [][]rune{},err
				}
				text[index-1] = decimal
			case "bin":
				decimal,err := Hex(text[index-1])
				if err != nil {
					return [][]rune{},err
				}
				text[index-1] = decimal
			case "up":
				if number == 0{
					word := Up(text[index-1])
					text[index-1] = word
				}else{
					for i :=0; i < number; i++{
						word := Up(text[index-i])
						text[index-i] = word
					}
				}
			case "low":
				if number == 0{
					word := Low(text[index-1])
					text[index-1] = word
				}else{
					for i :=0; i < number; i++{
						word := Low(text[index-i])
						text[index-i] = word
					}
				}
			case "cap":
				if number == 0{
					word := Cap(text[index-1])
					text[index-1] = word
				}else{
					for i :=0; i < number; i++{
						word := Cap(text[index-i])
						text[index-i] = word
					}
				}
			}
		}

	}
	return RemoveParameters(text), nil
}

func Parameters(parameter []rune)(string,int,error){
	if parameter[0] != '('{
		return "",0,errors.New("Ce qui est rentrer en paramètre n'est pas un paramètre")
	}
	word := ""
	number := 0
	isNumber := false
	for _,character := range parameter {
		if character == ')'{
			return word,number, nil
		}
		if isNumber == false && character != ','{
			if character != '('{
				word += string(character)
			}
		}else if isNumber == true{
			number = number*10 + int(character - '0')
		}
		if isNumber == false && character == ','{
			isNumber = true
		}
	}
	return "",0,errors.New("Sortie du paramètre sans trouver de parenthèse fermante")
}

func RemoveParameters(text [][]rune)([][]rune){
	newText := [][]rune{}
	for _,word := range text{
		if word[0] != '('{
			newText = append(newText,word)
		}
	}
	return newText
}