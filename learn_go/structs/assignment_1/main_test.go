package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		membershipType membershipType
	}{
		{"Syl", TypeStandard},
		{"Pattern", TypePremium},
		{"Pattern", TypeStandard},
	}
	if withSubmit {
		submitCases := []struct {
			name           string
			membershipType membershipType
		}{
			{"Renarin", TypeStandard},
			{"Lift", TypePremium},
			{"Dalinar", TypeStandard},
		}
		tests = append(tests, submitCases...)
	}

	for _, tc := range tests {
		user := newUser(tc.name, tc.membershipType)

		msgCharLimit := 100
		if tc.membershipType == TypePremium {
			msgCharLimit = 1000
		}

		if user.Name != tc.name {
			t.Errorf(`Test Failed (name):
  exected: %v
  actual: %s
`,
				tc.name, user.Name)
		} else if user.Type != tc.membershipType {
			t.Errorf(`Test Failed (membership type):
  exected: %v
  actual: %v
		`,
				tc.membershipType, user.Type)
		} else if user.MessageCharLimit != msgCharLimit {
			t.Errorf(`Test Failed (message character limit):
  exected: %v
  actual: %v
			`,
				msgCharLimit, user.MessageCharLimit)
		} else {
			fmt.Printf(`Test Passed:
* user: %s
* membership type: %s
* message character limit: %v
`,
				user.Name, user.Type, user.MessageCharLimit)
		}
		fmt.Println("==============================")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
