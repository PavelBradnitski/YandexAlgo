package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	input, _ := in.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	var values []string
	currentNumber := ""
	// for i, char := range input {
	// 	if char >= '0' && char <= '9' {
	// 		currentNumber += string(char)
	// 	} else {
	// 		if i == 0 || (i > 0 && input[i-1] == '(') {
	// 			if char == '-' {
	// 				values = append(values, "0")
	// 			}
	// 		}
	// 		if currentNumber != "" {
	// 			values = append(values, currentNumber)
	// 			currentNumber = ""
	// 		}
	// 		values = append(values, string(char))
	// 	}

	// 	if i > 0 && char >= '0' && char <= '9' && input[i-1] >= '0' && input[i-1] <= '9' {
	// 		fmt.Fprintln(out, "WRONG1")
	// 		return
	// 	}

	// }
	for i, char := range input {
		if char >= '0' && char <= '9' { // Проверка на цифру
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				values = append(values, currentNumber)
				currentNumber = ""
			}
			values = append(values, string(char))
		}
		// Проверка на два числа подряд (сразу после добавления цифры)
		if i > 0 && input[i-1] >= '0' && input[i-1] <= '9' {
			fmt.Fprintln(out, "WRONG1")
			return // Прерываем программу, если обнаружена ошибка
		}
		// Проверка на отрицательное число в начале строки или после открывающейся скобки
		if i == 0 || (i > 0 && input[i-1] == '(') {
			if char == '-' {
				values = append(values, "0") // Добавляем 0 перед отрицательным числом
			}
		}
	}
	// fmt.Println("Значения в срезе:", values)
	if currentNumber != "" {
		values = append(values, currentNumber)
	}
	output, err := Calculator(values)
	if err == nil {
		fmt.Fprintln(out, output)
	} else {
		fmt.Fprintln(out, "WRONG")
	}

}

func Calculator(str []string) (int, error) {
	var postfix []string
	var stack []rune
	prevSymb := string(str[0])
	if !RightBracketSeq(str) {
		return 0, fmt.Errorf("bracket error")
	}
	for i, v := range str {

		if i != 0 {
			_, err1 := strconv.Atoi(prevSymb)
			_, err2 := strconv.Atoi(string(v))
			if err1 == nil && err2 == nil {
				return 0, fmt.Errorf("can't be 2 digit in series")
			}
		}
		if _, err := strconv.Atoi(string(v)); err == nil {
			postfix = append(postfix, v)
			prevSymb = string(v)
		} else {
			switch v {
			case "+":
				for len(stack) != 0 && (stack[len(stack)-1] == '+' || stack[len(stack)-1] == '-') {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, '+')
				prevSymb = string(v)
			case "*":
				for len(stack) != 0 && stack[len(stack)-1] == '*' {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, '*')
				prevSymb = string(v)
			case "-":
				for len(stack) != 0 && (stack[len(stack)-1] == '+' || stack[len(stack)-1] == '-' || stack[len(stack)-1] == '*') {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}

				stack = append(stack, '-')
				prevSymb = string(v)
			case "(":
				stack = append(stack, '(')
				prevSymb = string(v)
			case ")":
				for len(stack) != 0 && stack[len(stack)-1] != '(' {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				if len(stack) != 0 && stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				}
				prevSymb = string(v)
			default:
				return 0, fmt.Errorf("unknown symbol")
			}
		}
	}
	for len(stack) != 0 {
		postfix = append(postfix, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}
	if containsRune(postfix, "(") {
		return 0, fmt.Errorf("error on postfix")
	}
	// for _, v := range postfix {
	// 	fmt.Printf("%s ", string(v))
	// }
	// fmt.Println("End Here")
	return CalculatePrefix(postfix)
}

func containsRune(runes []string, target string) bool {
	for _, r := range runes {
		if r == target {
			return true
		}
	}
	return false
}
func CalculatePrefix(inpStr []string) (int, error) {
	var stack []int
	for _, char := range inpStr {
		if num, err := strconv.Atoi(string(char)); err == nil {
			stack = append(stack, num)
		} else {
			s := stack[len(stack)-1]
			f := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch char {
			case "+":
				stack = append(stack, s+f)
			case "*":
				stack = append(stack, s*f)
			case "-":
				stack = append(stack, f-s)
			}
		}
	}
	if len(stack) > 1 {
		return 0, fmt.Errorf("invalid postfix")
	}
	return stack[0], nil
}

func RightBracketSeq(str []string) bool {
	var stack []string
	for _, v := range str {
		if v == "(" || v == "{" || v == "[" {
			stack = append(stack, v)
		} else if len(stack) > 0 {
			switch v {
			case ")":
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case "}":
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case "]":
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}
