package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var result []rune
	input, _ := in.ReadString('\n')
	for _, char := range input {
		if !unicode.IsSpace(char) {
			result = append(result, char)
		}
	}
	fmt.Fprintln(out, Calculate(result))

}

func Calculate(inpStr []rune) int {
	var stack []int
	for _, char := range inpStr {
		if num, err := strconv.Atoi(string(char)); err == nil {
			stack = append(stack, num)
		} else {
			s := stack[len(stack)-1]
			f := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch char {
			case '+':
				stack = append(stack, s+f)
			case '*':
				stack = append(stack, s*f)
			case '-':
				stack = append(stack, f-s)
			}
		}
	}
	return stack[0]
}
