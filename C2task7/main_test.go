package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expressionNum  int
		expectedResult int
	}{
		{
			name:           "Basic Test 1",
			expression:     "aab",
			expressionNum:  1,
			expectedResult: 2,
		},
		{
			name:           "Basic Test 2",
			expression:     "aabcbb",
			expressionNum:  2,
			expectedResult: 4,
		},
		{
			name:           "With 0 a + b",
			expression:     "sssss",
			expressionNum:  10,
			expectedResult: 5,
		},
		{
			name:           "With 0 a + b",
			expression:     "aaabaababaabaaaabaabbxaazbaaaababaababbbaaaabaabbb",
			expressionNum:  159,
			expectedResult: 39,
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val := Censor(testCase.expression, testCase.expressionNum)
			if val != testCase.expectedResult {
				t.Fatalf("%d should be equal %d", val, testCase.expectedResult)
			}
		})
	}
}
