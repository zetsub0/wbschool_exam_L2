package main

import "testing"

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}
	for _, testCase := range testCases {
		res, err := Unpack(testCase.input)
		if err != nil {
			if testCase.expected != "" {
				t.Errorf("Expected no error, but got error: %v", err)
			}
		} else if res != testCase.expected {
			t.Errorf("For input \"%s\", expected \"%s\", but got \"%s\"", testCase.input, testCase.expected, res)
		}
	}
}
