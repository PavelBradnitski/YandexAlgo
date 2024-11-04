package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     []int
		expectedResult string
	}{
		{
			name:           "Basic Test 1",
			expression:     []int{19, 8, 3},
			expectedResult: "8 3 19",
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val := Median(testCase.expression)
			if val != testCase.expectedResult {
				t.Fatalf("%s should be equal %s", val, testCase.expectedResult)
			}
		})
	}
}
