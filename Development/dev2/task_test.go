package dev2

import "testing"

func TestUnpackGeneralPositiveCheck(t *testing.T) {
	example := "a4bb8c7aa"
	must := "aaaabbbbbbbbbcccccccaa"

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); result != must || err != nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", must, result)
	}
}

func TestString_Unpack_UnicodePositiveCheck(t *testing.T) {
	example := "a4b假8c7aa"
	must := "aaaab假假假假假假假假cccccccaa"

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); result != must || err != nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", must, result)
	}
}

func TestString_Unpack_IncorrectInputNegativeCheck(t *testing.T) {
	example := "35"

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); err == nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", "error", result)
	}
}

func TestString_Unpack_EmptyInputPositiveCheck(t *testing.T) {
	example := ""
	must := ""

	var ext Extractor = &String{}

	if result, err := ext.Unpack(example); result != must || err != nil {
		t.Fatalf("Wrong result: want (%s), got (%s).", must, result)
	}
}

//---------
// package main

// import (
// 	"testing"
// )

// func TestUnpackString(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected string
// 		hasError bool
// 	}{
// 		{"a4bc2d5e", "aaaabccddddde", false},
// 		{"abcd", "abcd", false},
// 		{"45", "", true},
// 		{"", "", false},
// 		{"qwe\\4\\5", "qwe45", false},
// 		{"qwe\\45", "qwe44444", false},
// 		{"qwe\\\\5", "qwe\\\\\\", false},
// 	}

// 	for _, test := range tests {
// 		result, err := UnpackString(test.input)
// 		if test.hasError {
// 			if err == nil {
// 				t.Errorf("expected error for input %q", test.input)
// 			}
// 		} else {
// 			if err != nil {
// 				t.Errorf("unexpected error for input %q: %v", test.input, err)
// 			}
// 			if result != test.expected {
// 				t.Errorf("for input %q expected %q but got %q", test.input, test.expected, result)
// 			}
// 		}
// 	}
// }
