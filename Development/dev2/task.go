package dev2

import (
	"errors"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Extractor interface {
	Unpack(s string) (string, error)
}

type String struct {
}

func (s *String) Unpack(str string) (string, error) {
	runes := []rune(str)
	answer := strings.Builder{}

	var esc = false
	var prev rune

	for _, v := range runes {
		switch {
		case esc:
			answer.WriteRune(v)
			prev = v
			esc = false
			continue
		case unicode.IsDigit(v):
			if prev == 0 {
				return answer.String(), errors.New("incorrect input string")
			}

			c := int(v) - '0'
			for i := 0; i < c-1; i++ {
				answer.WriteRune(prev)
			}
			continue
		case v == '/':
			esc = true
			continue
		default:
			prev = v
			answer.WriteRune(v)
		}
	}

	return answer.String(), nil
}

//---
// package main

// import (
// 	"errors"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// // UnpackString распаковывает строку согласно заданному правилу
// func UnpackString(input string) (string, error) {
// 	var result strings.Builder
// 	runes := []rune(input)
// 	escape := false

// 	for i := 0; i < len(runes); i++ {
// 		switch {
// 		case escape:
// 			result.WriteRune(runes[i])
// 			escape = false
// 		case runes[i] == '\\':
// 			escape = true
// 		case unicode.IsDigit(runes[i]):
// 			if i == 0 || unicode.IsDigit(runes[i-1]) {
// 				return "", errors.New("invalid string")
// 			}
// 			count, _ := strconv.Atoi(string(runes[i]))
// 			result.WriteString(strings.Repeat(string(runes[i-1]), count-1))
// 		default:
// 			result.WriteRune(runes[i])
// 		}
// 	}
// 	if escape {
// 		return "", errors.New("invalid string")
// 	}
// 	return result.String(), nil
// }
