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
			name:           "Go to North",
			expression:     []int{-1, -2, 5, 3, -1, 6},
			expectedResult: "N",
		},
		{
			name:           "Go to South",
			expression:     []int{-1, -2, 5, 3, -1, -3},
			expectedResult: "S",
		},
		{
			name:           "Go to East",
			expression:     []int{-1, -2, 5, 3, 6, -1},
			expectedResult: "E",
		},
		{
			name:           "Go to West",
			expression:     []int{-1, -2, 5, 3, -2, -1},
			expectedResult: "W",
		},
		{
			name:           "Go to NorthEast",
			expression:     []int{-1, -2, 5, 3, 6, 6},
			expectedResult: "NE",
		},
		{
			name:           "Go to SouthEast",
			expression:     []int{-1, -2, 5, 3, 6, -3},
			expectedResult: "SE",
		},
		{
			name:           "Go to NorthWest",
			expression:     []int{-1, -2, 5, 3, -3, 6},
			expectedResult: "NW",
		},
		{
			name:           "Go to SouthWest",
			expression:     []int{-1, -2, 5, 3, -6, -3},
			expectedResult: "SW",
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val := DefinePath(testCase.expression[0], testCase.expression[1], testCase.expression[2], testCase.expression[3], testCase.expression[4], testCase.expression[5])
			if val != testCase.expectedResult {
				t.Fatalf("%s should be equal %s", val, testCase.expectedResult)
			}
		})
	}
}
