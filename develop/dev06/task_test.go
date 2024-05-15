package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCut struct {
	f, d           string
	s              bool
	data           string
	expectedString string
	expectedError  error
}

var testCuts = []testCut{
	{
		"1,2,3",
		" ",
		false,
		"First_Test_First_Line\nFirst test Second Line\nFirstTestThirdLine",
		"First_Test_First_Line\nFirst test Second\nFirstTestThirdLine",
		nil,
	},
	{
		"1,2",
		" ",
		false,
		"Second_Test_First_Line\nSecond Test Second Line\nSecond_Test Third Line",
		"Second_Test_First_Line\nSecond Test\nSecond_Test Third",
		nil,
	},
	{
		"3,4",
		"_",
		true,
		"Third_Test_First_Line\nThird Test Second Line\nThirdTestThird Line",
		"First_Line",
		nil,
	},
}

func TestCut(t *testing.T) {
	for _, test := range testCuts {
		output, err := cut(test.data, test.f, test.d, test.s)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedString, output)
	}
}
