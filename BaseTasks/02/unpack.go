package main

import (
	"errors"
	"fmt"
	"unicode"
)

//Создать Go-функцию, осуществляющую примитивную распаковку
//строки, содержащую повторяющиеся символы/руны
func main() {
	s, _ := Unpack("a4bc2d5e")
	s, _ = Unpack("45")
	fmt.Println(s)
}

func Unpack(s string) (result string, err error) {
	r := []rune(s)
	var previous rune
	for i, sym := range r {
		if unicode.IsDigit(sym) {
			if i == 0 || unicode.IsDigit(previous) {
				err = errors.New("unpack: Incorrect Input")
				return
			} else {
				result += repeat(previous, sym)
			}
		} else {
			if i != 0 && !unicode.IsDigit(previous) {
				result += repeat(previous, '1')
			}
		}
		previous = sym
		if len(r)-1 == i && !unicode.IsDigit(previous) {
			result += repeat(previous, '1')
		}
	}
	return
}

func repeat(r rune, i rune) string {
	var resultR []rune
	for j := '0'; j < i; j++ {
		resultR = append(resultR, r)
	}
	return string(resultR)
}
