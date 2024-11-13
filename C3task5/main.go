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

	var values []string
	currentNumber := ""
	var err error
	for i, char := range input {
		if char != '(' && char != ')' && char != '+' && char != '*' && char != '-' && char != ' ' {
			if char < '0' && char > '9' {
				fmt.Fprintln(out, "WRONG")
				return
			}
		}
		if char >= '0' && char <= '9' {
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				values = append(values, currentNumber)
				currentNumber = ""
			}
			if i == 0 || (i > 0 && input[i-1] == '(') {
				if char == '-' {
					values = append(values, "0")
				}
			}
			if i == 0 || (i > 0 && input[i-1] == '(') {
				if char == ')' || char == '+' || char == '*' {
					fmt.Fprintln(out, "WRONG")
					return
				}
			}
			if char != ' ' {
				values = append(values, string(char))
			}
		}
		if len(values) != 0 {
			_, err = strconv.Atoi(values[len(values)-1])
		}
		if len(values) != 0 && currentNumber != "" && err == nil {
			fmt.Fprintln(out, "WRONG")
			return
		}

	}
	if currentNumber != "" {
		values = append(values, currentNumber)
	}
	if output, err := Calculator(values); err == nil {
		fmt.Fprintln(out, output)
	} else {
		fmt.Fprintln(out, "WRONG")
	}
}

func Calculator(str []string) (int64, error) {
	var postfix []string
	var stack []rune
	if !RightBracketSeq(str) {
		return 0, fmt.Errorf("bracket error")
	}
	for _, v := range str {
		if _, err := strconv.Atoi(string(v)); err == nil {
			postfix = append(postfix, v)
		} else {
			switch v {
			case "+":
				for len(stack) != 0 && (stack[len(stack)-1] == '+' || stack[len(stack)-1] == '-' || stack[len(stack)-1] == '*') {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, '+')
			case "*":
				for len(stack) != 0 && stack[len(stack)-1] == '*' {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, '*')
			case "-":
				for len(stack) != 0 && (stack[len(stack)-1] == '+' || stack[len(stack)-1] == '-' || stack[len(stack)-1] == '*') {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, '-')
			case "(":
				stack = append(stack, '(')
			case ")":
				for len(stack) != 0 && stack[len(stack)-1] != '(' {
					postfix = append(postfix, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				if len(stack) != 0 && stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				}
			default:
				return 0, fmt.Errorf("unknown symbol")
			}
		}
	}
	for len(stack) != 0 {
		postfix = append(postfix, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}
	return CalculatePrefix(postfix)
}

func CalculatePrefix(inpStr []string) (int64, error) {
	var stack []int64
	var s, f int64
	for _, char := range inpStr {
		if num, err := strconv.ParseInt(string(char), 10, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) > 1 {
				s = stack[len(stack)-1]
				f = stack[len(stack)-2]
				stack = stack[:len(stack)-2]
			}
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
		if v == "(" {
			stack = append(stack, v)
		} else if len(stack) > 0 {
			switch v {
			case ")":
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				}
			}
		} else if v == ")" {
			return false
		}
	}
	return len(stack) == 0
}
