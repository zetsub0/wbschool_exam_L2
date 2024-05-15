package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySort(t *testing.T) {
	tests := []struct {
		r, u, n       bool
		k             int
		data          []byte
		expected      string
		expectedError bool
	}{
		{
			r:             false,
			u:             false,
			n:             false,
			k:             0,
			data:          []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected:      "aaaa pfor 9874\nbbbb olws 6714\ndddd ebcd 21234\nnnnn yhji 2912\noooo ikik 9120\nssss odod 2921\nuuuu vfjv 8212\n",
			expectedError: false,
		},
		{
			r:             false,
			u:             false,
			n:             true,
			k:             0,
			data:          []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected:      "2 ssss odod 2921\n5 uuuu vfjv 8212\n6 oooo ikik 9120\n10 dddd ebcd 12343\n32 nnnn yhji 2912\n50 bbbb olws 6714\n100 aaaa pfor 9874\n",
			expectedError: false,
		},
		{
			r:             true,
			u:             false,
			n:             false,
			k:             0,
			data:          []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected:      "6 oooo ikik 9120\n5 uuuu vfjv 8212\n50 bbbb olws 6714\n32 nnnn yhji 2912\n2 ssss odod 2921\n10 dddd ebcd 12343\n100 aaaa pfor 9874\n",
			expectedError: false,
		},
		{
			r:             false,
			u:             false,
			n:             false,
			k:             4,
			data:          []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected:      "10 dddd ebcd 12343\n32 nnnn yhji 2912\n2 ssss odod 2921\n50 bbbb olws 6714\n5 uuuu vfjv 8212\n6 oooo ikik 9120\n100 aaaa pfor 9874\n",
			expectedError: false,
		},
		{
			r:             false,
			u:             false,
			n:             true,
			k:             4,
			data:          []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected:      "32 nnnn yhji 2912\n2 ssss odod 2921\n50 bbbb olws 6714\n5 uuuu vfjv 8212\n6 oooo ikik 9120\n100 aaaa pfor 9874\n10 dddd ebcd 12343\n",
			expectedError: false,
		},
	}

	for _, test := range tests {
		output, _ := mySort(test.data, test.r, test.n, test.u, test.k)
		if test.expectedError {
			assert.Equal(t, test.expected, output)
		}
	}
}
