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
	var len, num int
	fmt.Fscan(in, &len, &num)
	inputStr := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Fscan(in, &inputStr[i])
	}
	fmt.Fprint(out, CountMonument(inputStr, num))

}
func CountMonument(nums []int, num int) int {
	n := len(nums)
	out := 0
	for i, j := 0, 1; j < n; {
		if nums[j]-nums[i] > num {
			out += n - j
			i++
		} else {
			j++
		}
	}
	return out
}

// func PrefixSum(nums []int) []int {
// 	numsToStr := make([]int, len(nums)+1)
// 	for i, v := range nums {
// 		numsToStr[i+1] = numsToStr[i] + v
// 	}
// 	return numsToStr
// }
