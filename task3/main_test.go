package main

import (
	"testing"
)

func TestFigure(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     []string
		func1          string
		expectedResult bool
	}{
		{
			name:           "Test I 1",
			expression:     []string{".###.", ".###.", ".###.", ".###."},
			func1:          "CheckI",
			expectedResult: true,
		},
		{
			name:           "Test I 1",
			expression:     []string{".##.", ".##.", ".##.", "...."},
			func1:          "CheckI",
			expectedResult: true,
		},
		{
			name:           "Test I Neg 1",
			expression:     []string{"###.", "###.", ".##.", "...."},
			func1:          "CheckI",
			expectedResult: false,
		},
		{
			name:           "Test I Neg 2",
			expression:     []string{".##.", "###.", ".##.", "####"},
			func1:          "CheckI",
			expectedResult: false,
		},
		{
			name:           "Test I Neg 3",
			expression:     []string{".###.", ".#.#.", ".###.", ".###."},
			func1:          "CheckI",
			expectedResult: false,
		},
		{
			name:           "Test O Pos 1",
			expression:     []string{".####.", ".#..#.", ".#..#.", ".####."},
			func1:          "CheckO",
			expectedResult: true,
		},
		{
			name: "Test O Pos 2",
			expression: []string{".#####.",
				".##..#.",
				".##..#.",
				".#####."},
			func1:          "CheckO",
			expectedResult: true,
		},
		{
			name:           "Test O Neg 1",
			expression:     []string{".#####.", ".#...#.", ".##..#.", ".#####."},
			func1:          "CheckO",
			expectedResult: false,
		},
		{
			name:           "Test O Neg 2",
			expression:     []string{".#####.", ".#.....", ".#.....", ".#####."},
			func1:          "CheckO",
			expectedResult: false,
		},
		{
			name:           "Test O Neg 3",
			expression:     []string{".#####.", ".#...#.", ".#...#.", ".##.##."},
			func1:          "CheckO",
			expectedResult: false,
		},
		{
			name:           "Test O Neg 4",
			expression:     []string{".#####.", ".#...#.", ".#...#.", ".##.##."},
			func1:          "CheckO",
			expectedResult: false,
		},
		{
			name:           "Test C Pos 1",
			expression:     []string{".#####.", ".##....", ".##....", ".#####."},
			func1:          "CheckC",
			expectedResult: true,
		},
		{
			name:           "Test C Pos 2",
			expression:     []string{".#####.", ".#.....", ".#.....", ".#####."},
			func1:          "CheckC",
			expectedResult: true,
		},
		{
			name:           "Test C Neg 1",
			expression:     []string{".#####.", ".#.....", ".#.....", ".####.."},
			func1:          "CheckC",
			expectedResult: false,
		},
		{
			name:           "Test C Neg 2",
			expression:     []string{"..####.", ".#.....", ".#.....", ".####.."},
			func1:          "CheckC",
			expectedResult: false,
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.func1 == "CheckI" {
				val := CheckI(testCase.expression)
				if val != testCase.expectedResult {
					t.Fatalf("%t should be equal %t", val, testCase.expectedResult)
				}
			} else if testCase.func1 == "CheckO" {
				val := CheckO(testCase.expression)
				if val != testCase.expectedResult {
					t.Fatalf("%t should be equal %t", val, testCase.expectedResult)
				}
			}

		})
	}
}
