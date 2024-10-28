package main

import (
	"fmt"
)

// Определим типы строк для упрощенного анализа
const (
	All  = "#"   // Строка целиком заполнена `#`
	Edge = "#.#" // Включены крайние элементы строки
	Left = "#"   // Включен только один крайний элемент строки
	None = "."   // Все элементы строки выключены
)

// Определяем букву на табло
func identifyLetter(matrix []string) string {
	// Обрезаем матрицу по границам символа
	trimmed := trimMatrix(matrix)
	// Строим представление для каждой строки
	lineTypes := analyzeLines(trimmed)

	// Определяем букву по анализу типов строк
	switch {
	case isI(lineTypes):
		return "I"
	case isO(lineTypes):
		return "O"
	case isC(lineTypes):
		return "C"
	case isL(lineTypes):
		return "L"
	case isH(lineTypes):
		return "H"
	case isP(lineTypes):
		return "P"
	default:
		return "X"
	}
}

// Обрезает матрицу по внешним границам символа
func trimMatrix(matrix []string) []string {
	top, bottom, left, right := -1, -1, -1, -1

	// Ищем крайние координаты включенных элементов
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '#' {
				if top == -1 || i < top {
					top = i
				}
				if bottom == -1 || i > bottom {
					bottom = i
				}
				if left == -1 || j < left {
					left = j
				}
				if right == -1 || j > right {
					right = j
				}
			}
		}
	}

	// Создаем обрезанную матрицу
	trimmed := make([]string, bottom-top+1)
	for i := top; i <= bottom; i++ {
		trimmed[i-top] = matrix[i][left : right+1]
	}
	return trimmed
}

// Анализирует каждую строку матрицы и определяет её тип
func analyzeLines(matrix []string) []string {
	var lineTypes []string
	for _, row := range matrix {
		lineTypes = append(lineTypes, analyzeLine(row))
	}
	return lineTypes
}

// Определяет тип строки (All, Edge, Left или None)
func analyzeLine(row string) string {
	if isAll(row) {
		return All
	} else if isEdge(row) {
		return Edge
	} else if isLeft(row) {
		return Left
	} else {
		return None
	}
}

// Вспомогательные функции для проверки типа строки
func isAll(row string) bool {
	for _, val := range row {
		if val != '#' {
			return false
		}
	}
	return true
}

func isEdge(row string) bool {
	return row[0] == '#' && row[len(row)-1] == '#' && allDots(row[1:len(row)-1])
}

func isLeft(row string) bool {
	return row[0] == '#' && allDots(row[1:])
}

func allDots(row string) bool {
	for _, val := range row {
		if val != '.' {
			return false
		}
	}
	return true
}

// Проверки на основе последовательности типов строк

func isI(lineTypes []string) bool {
	// Вся матрица должна быть типа "All"
	for _, line := range lineTypes {
		if line != All {
			return false
		}
	}
	return true
}

func isO(lineTypes []string) bool {
	// Верхняя и нижняя строки — Edge, внутри — None
	return lineTypes[0] == Edge && lineTypes[len(lineTypes)-1] == Edge && all(lineTypes[1:len(lineTypes)-1], None)
}

func isC(lineTypes []string) bool {
	// Аналогично `O`, но правая граница — Edge
	return lineTypes[0] == Edge && all(lineTypes[1:], Edge)
}

func isL(lineTypes []string) bool {
	// Верхняя граница — Edge, остальное — All
	return lineTypes[0] == Edge && all(lineTypes[1:], All)
}

func isH(lineTypes []string) bool {
	// Ищем две крайние Edge с блоком None между ними
	top, bottom := lineTypes[0] == Edge, lineTypes[len(lineTypes)-1] == Edge
	return top && bottom && all(lineTypes[1:len(lineTypes)-1], None)
}

func isP(lineTypes []string) bool {
	// Верх и низ Edge, средний — None, правый — Edge
	return lineTypes[0] == Edge && all(lineTypes[1:len(lineTypes)-1], None) && lineTypes[len(lineTypes)-1] == Edge
}

func all(slice []string, val string) bool {
	for _, item := range slice {
		if item != val {
			return false
		}
	}
	return true
}

// Пример использования
func main() {
	var len int

	fmt.Scanln(&len)
	inputStr := make([]string, len)
	for i := 0; i < len; i++ {
		fmt.Scanln(&inputStr[i])
	}

	//fmt.Println(inputStr)

	fmt.Println(identifyLetter(inputStr))
}
