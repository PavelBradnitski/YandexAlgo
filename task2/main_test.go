package main

import (
	"testing"
)

func TestFindMinComb(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     []int
		expectedResult []int
	}{
		{
			name:           "test6 ",
			expression:     []int{10, 7, 0, 4},
			expectedResult: []int{11, 1},
		},
		{
			name:           "test4 ",
			expression:     []int{0, 2, 5, 1},
			expectedResult: []int{1, 6},
		},
		{
			name:           "test25 ",
			expression:     []int{114, 299, 921, 166},
			expectedResult: []int{300, 1},
		},
		{
			name:           "test26 ",
			expression:     []int{790, 493, 507, 302},
			expectedResult: []int{1, 508},
		},
		{
			name:           "Base test",
			expression:     []int{6, 2, 7, 3},
			expectedResult: []int{3, 4},
		},
		{
			name:           "Revert Base test",
			expression:     []int{2, 6, 3, 7},
			expectedResult: []int{3, 4},
		},

		{
			name:           "Small numbers",
			expression:     []int{2, 1, 1, 1},
			expectedResult: []int{1, 2},
		},
		{
			name:           "Small numbers",
			expression:     []int{3, 1, 2, 2},
			expectedResult: []int{1, 3},
		},
		{
			name:           "Zero case",
			expression:     []int{3, 0, 5, 2},
			expectedResult: []int{1, 3},
		},
		{
			name:           "Zero2 case",
			expression:     []int{3, 0, 5, 0},
			expectedResult: []int{1, 1},
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val1, val2 := FindMinComb(testCase.expression[0], testCase.expression[1], testCase.expression[2], testCase.expression[3])
			if val1 != testCase.expectedResult[0] || val2 != testCase.expectedResult[1] {
				t.Fatalf("%d + %d should be equal %v", val1, val2, testCase.expectedResult)
			}
		})
	}
}
