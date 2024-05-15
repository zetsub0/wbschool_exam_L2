package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type grepTest struct {
	A, B, C        int
	c, i, v, F, n  bool
	data, pattern  string
	expectedString string
	expectedError  error
}

var grepTests = []grepTest{
	{
		0,
		0,
		0,
		false,
		false,
		false,
		false,
		false,
		"first\nsecond\nHelp\nhelp\nHelp\nh3lp\nHeeelp\nHide and seek\nJeromo\nKurd\nCringe\nHello world\nDefer pani\nlog fatal",
		"Help", "Help\nHelp\n",
		nil,
	},
	{
		0,
		0,
		0,
		false,
		true,
		false,
		false,
		false,
		"first\nsecond\nHelp\nhelp\nHelp\nh3lp\nHeeelp\nHide and seek\nJeromo\nKurd\nCringe\nHello world\nDefer pani\nlog fatal",
		"HELP", "Help\nhelp\nHelp\n",
		nil,
	},
	{
		2,
		0,
		0,
		true,
		true,
		false,
		false,
		true,
		"first\nsecond\nHelp\nhelp\nHelp\nh3lp\nHeeelp\nHide and seek\nJeromo\nKurd\nCringe\nHello world\nDefer pani\nlog fatal",
		"HeLp",
		"3",
		nil,
	},
	{
		0,
		0,
		2,
		false,
		false,
		false,
		false,
		false,
		"first\nsecond\nHelp\nhelp\nHelp\nh3lp\nHeeelp\nHide and seek\nJeromo\nKurd\nCringe\nHello world\nDefer pani\nlog fatal",
		"seek",
		"h3lp\nHeeelp\nHide and seek\nJeromo\nKurd\n",
		nil,
	},
	{
		0,
		0,
		0,
		true,
		true,
		true,
		false,
		true,
		"first\nsecond\nHelp\nhelp\nHelp\nh3lp\nHeeelp\nHide and seek\nJeromo\nKurd\nCringe\nHello world\nDefer pani\nlog fatal",
		"Crin",
		"13",
		nil,
	},
}

func TestGrep(t *testing.T) {
	for _, test := range grepTests {
		output, err := grep(test.data, test.pattern, test.A, test.B, test.C, test.c, test.i, test.v, test.F, test.n)
		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedString, output)
	}
}
