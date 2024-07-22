
package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name     string
		expected uintptr
	}

	tests := []testCase{
		{"contact", uintptr(24)},
	}
	if withSubmit {
		tests = append(tests, testCase{"perms", uintptr(16)})
	}

	for _, test := range tests {
		var typ reflect.Type
		if test.name == "contact" {
			typ = reflect.TypeOf(contact{})
		} else if test.name == "perms" {
			typ = reflect.TypeOf(perms{})
		}

		size := typ.Size()

		if size != test.expected {
			t.Errorf(`Test Failed: (%v)
  expected: %v bytes
  actual: %v bytes
`,
				test.name, test.expected, size)
		} else {
			fmt.Printf(`Test Passed: (%v)
  expected: %v bytes
  actual: %v bytes
`,
				test.name, test.expected, size)
		}
		fmt.Println("----------------")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
