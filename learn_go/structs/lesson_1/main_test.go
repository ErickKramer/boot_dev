package main

import (
	"fmt"
	"testing"
)

func getMessageText(m messageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
}

func Test(t *testing.T) {
	type testCase struct {
		phoneNumber int
		message     string
		expected    string
	}
	tests := []testCase{
		{148255510981, "Thanks for signing up", "Sending message: 'Thanks for signing up' to: 148255510981"},
		{148255510982, "Love to have you aboard!", "Sending message: 'Love to have you aboard!' to: 148255510982"},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{148255510983, "We're so excited to have you", "Sending message: 'We're so excited to have you' to: 148255510983"},
			{148255510984, "", "Sending message: '' to: 148255510984"},
			{148255510985, "Hello, World!", "Sending message: 'Hello, World!' to: 148255510985"},
		}...)
	}

	for _, test := range tests {
		if output := getMessageText(messageToSend{
			phoneNumber: test.phoneNumber,
			message:     test.message,
		}); output != test.expected {
			t.Errorf(
				`Test Failed: (%v, %v) ->
  expected: %v
  actual: %v
`,
				test.phoneNumber,
				test.message,
				test.expected,
				output,
			)
		} else {
			fmt.Printf(`Test Passed: (%v, %v) ->
  expected: %v
  actual: %v
`,
				test.phoneNumber,
				test.message,
				test.expected,
				output,
			)
		}
		fmt.Println("================================")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
