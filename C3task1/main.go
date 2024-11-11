package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var inputStr string
	fmt.Fscan(in, &inputStr)
	if RightBracketSeq(inputStr) {
		fmt.Fprint(out, "yes")
	} else {
		fmt.Fprint(out, "no")
	}
}

func RightBracketSeq(str string) bool {
	var stack []rune
	for _, v := range str {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if len(stack) > 0 {
			switch v {
			case ')':
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case '}':
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			case ']':
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}
