package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     []int
		expressionNum  int
		expectedResult int
	}{
		{
			name:           "Basic Test 1",
			expression:     []int{4, 2, 1},
			expressionNum:  2,
			expectedResult: 2,
		},
		{
			name:           "Basic Test 2",
			expression:     []int{3, 8, 5, 7, 1, 2, 4, 9, 6},
			expressionNum:  2,
			expectedResult: 3,
		},
		{
			name:           "My Test 1",
			expression:     []int{1, 3, 1},
			expressionNum:  0,
			expectedResult: 2,
		},
		{
			name:           "My Test 2",
			expression:     []int{32, 77, 1, 100},
			expressionNum:  4,
			expectedResult: 1,
		},
		{
			name:           "My Test 2",
			expression:     []int{1, 1, 1, 1},
			expressionNum:  0,
			expectedResult: 4,
		},
		{
			name:           "My Test 2",
			expression:     []int{14, 15, 15, 16, 16, 16, 16, 17, 17, 18},
			expressionNum:  2,
			expectedResult: 8,
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val := minPasses(testCase.expression, testCase.expressionNum)
			if val != testCase.expectedResult {
				t.Fatalf("%d should be equal %d", val, testCase.expectedResult)
			}
		})
	}
}
