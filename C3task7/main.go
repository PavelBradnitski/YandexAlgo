package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type InputStr struct {
	Operation rune
	Num       int64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var length int
	var str string
	fmt.Fscanln(in, &length)
	inputSlice := make([]InputStr, length)
	for i := 0; i < length; i++ {
		fmt.Fscanln(in, &str)
		inputSlice[i].Operation = rune(str[0])
		if inputSlice[i].Operation != '-' {
			inputSlice[i].Num, _ = strconv.ParseInt(str[1:], 10, 64)
		}

	}
	output := Stack(inputSlice)
	for _, v := range output {
		fmt.Fprintln(out, v)
	}

}
func Stack(inp []InputStr) []int64 {
	var out, stack, prefixSum []int64
	prefixSum = append(prefixSum, 0)
	for _, v := range inp {
		switch v.Operation {
		case '+':
			stack = append(stack, v.Num)
			prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+v.Num)
		case '-':
			out = append(out, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			prefixSum = prefixSum[:len(prefixSum)-1]
		case '?':
			out = append(out, prefixSum[len(prefixSum)-1]-prefixSum[len(prefixSum)-1-int(v.Num)])
		}
	}
	return out
}
