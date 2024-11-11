package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult bool
	}{
		{
			name:           "Basic Test 1",
			expression:     "()[]",
			expectedResult: true,
		},
		{
			name:           "Basic Test 2",
			expression:     "([)]",
			expectedResult: false,
		},
		{
			name:           "Basic Test 3",
			expression:     "(",
			expectedResult: false,
		},
		{
			name:           "My Test 1",
			expression:     "()()()",
			expectedResult: true,
		},
		{
			name:           "My Test 1",
			expression:     "(((((((((([[[[[[[[[[[[",
			expectedResult: false,
		},
		{
			name:           "My Test 1",
			expression:     "())(()",
			expectedResult: false,
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val := RightBracketSeq(testCase.expression)
			if val != testCase.expectedResult {
				t.Fatalf("%t should be equal %t", val, testCase.expectedResult)
			}
		})
	}
}
