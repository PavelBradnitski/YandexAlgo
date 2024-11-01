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
	var length int
	fmt.Fscanln(in, &length)
	inputStr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(in, &inputStr[i])
	}
	numbers := PrefixSum(inputStr)
	numbers = numbers[1:]
	for i := 0; i < len(numbers); i++ {
		fmt.Fprintf(out, "%d ", numbers[i])
	}

}
func PrefixSum(nums []int) []int {
	numsToStr := make([]int, len(nums)+1)
	for i, v := range nums {
		numsToStr[i+1] = numsToStr[i] + v
	}
	return numsToStr
}
