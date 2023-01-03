package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name   string
		cases  string
		result bool
	}{
		{
			name:   "When we use one Uppercase and other Lowercase letters",
			cases:  "Hello",
			result: true,
		}, {
			name:   "When we use only Uppercase letters",
			cases:  "HELLO",
			result: true,
		}, {
			name:   "When we use 2 letters: the first is Uppercase and the second Lowercase",
			cases:  "Hi",
			result: true,
		}, {
			name:   "When we use one Uppercase letter at the end and length of word == 3",
			cases:  "HeL",
			result: false,
		}, {
			name:   "When we use one Uppercase letter at the end and length of word > 3",
			cases:  "HellO",
			result: false,
		}, {
			name:   "When we use word which its contain no letter",
			cases:  "He5lo",
			result: false,
		}, {
			name:   "When we use more than one Uppercase letters and it has Lowercase",
			cases:  "HELlo",
			result: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, DetectCapitalUse(tc.cases), tc.result)
		})
	}
}
