package main

import (
	"testing"
)

func TestFigure(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     []string
		expectedResult string
	}{
		{
			name:           "Test I Pos 1",
			expression:     []string{".###.", ".###.", ".###.", ".###."},
			expectedResult: "I",
		},
		{
			name:           "Test I Pos 2",
			expression:     []string{".##.", ".##.", ".##.", "...."},
			expectedResult: "I",
		},
		{
			name:           "Test I Neg 1",
			expression:     []string{"###.", "###.", ".##.", "...."},
			expectedResult: "X",
		},
		{
			name:           "Test I Neg 2",
			expression:     []string{".##.", "###.", ".##.", "####"},
			expectedResult: "X",
		},
		{
			name:           "Test I Neg 3",
			expression:     []string{".###.", ".#.#.", ".###.", ".###."},
			expectedResult: "X",
		},
		{
			name:           "Test O Pos 1",
			expression:     []string{".####.", ".#..#.", ".#..#.", ".####."},
			expectedResult: "O",
		},
		{
			name: "Test O Pos 2",
			expression: []string{".#####.",
				".##..#.",
				".##..#.",
				".#####."},
			expectedResult: "O",
		},
		{
			name:           "Test O Neg 1",
			expression:     []string{".#####.", ".#...#.", ".##..#.", ".#####."},
			expectedResult: "X",
		},
		{
			name:           "Test O Neg 2",
			expression:     []string{".#####.", ".#.....", ".#.....", ".#####."},
			expectedResult: "X",
		},
		{
			name:           "Test O Neg 3",
			expression:     []string{".#####.", ".#...#.", ".#...#.", ".##.##."},
			expectedResult: "X",
		},
		{
			name:           "Test O Neg 4",
			expression:     []string{".#####.", ".#...#.", ".#...#.", ".##.##."},
			expectedResult: "X",
		},
		{
			name:           "Test C Pos 1",
			expression:     []string{".#####.", ".##....", ".##....", ".#####."},
			expectedResult: "C",
		},
		{
			name:           "Test C Pos 2",
			expression:     []string{".#####.", ".#.....", ".#.....", ".#####."},
			expectedResult: "C",
		},
		{
			name:           "Test C Neg 1",
			expression:     []string{".#####.", ".#.....", ".#.....", ".####.."},
			expectedResult: "X",
		},
		{
			name:           "Test C Neg 2",
			expression:     []string{"..####.", ".#.....", ".#.....", ".####.."},
			expectedResult: "C",
		},
	}
	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val := identifyLetter(testCase.expression)
			if val != testCase.expectedResult {
				t.Fatalf("%s should be equal %s", val, testCase.expectedResult)
			}
		})
	}
}
