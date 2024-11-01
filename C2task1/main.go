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
	var len int
	fmt.Fscanln(in, &len)
	inputStr := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Fscan(in, &inputStr[i])
	}
	fmt.Fprint(out, strings.Join(PrefixSum(inputStr), " "))

}
func PrefixSum(nums []int) []string {
	sum := 0
	numsToStr := make([]string, len(nums))
	for i, v := range nums {
		sum += v
		numsToStr[i] = strconv.Itoa(sum)
	}
	return numsToStr
}
