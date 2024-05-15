package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		data     []string
		expected map[string][]string
	}{
		{
			data:     []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "ток", "кот", "окт", "кто"},
			expected: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}, "ток": {"кот", "кто", "окт", "ток"}},
		},
		{
			data:     []string{"eat", "ate", "teA", "bike", "kibe", "cab", "bca", "Abc", "aBc", "a", "b"},
			expected: map[string][]string{"eat": {"ate", "eat", "tea"}, "bike": {"bike", "kibe"}, "cab": {"abc", "bca", "cab"}}},
		{
			data:     []string{"eat", "eat", "eat", "ate", "teA", "bike", "kibe", "cab", "bca", "Abc", "aBc"},
			expected: map[string][]string{"bike": []string{"bike", "kibe"}, "cab": []string{"abc", "bca", "cab"}, "eat": []string{"ate", "eat", "tea"}}},
	}

	for _, test := range tests {
		output := findAnagrams(test.data)
		assert.Equal(t, test.expected, output)
	}
}
