package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var length, num int
	fmt.Fscan(in, &length, &num)
	intSlice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(in, &intSlice[i])
	}
	fmt.Fprint(out, minPasses(intSlice, num))

}

func minPasses(numbers []int, n int) int {
	sort.Ints(numbers)
	var days, curGr, left, right int
	for right < len(numbers) {
		for right < len(numbers) && numbers[right]-numbers[left] <= n {
			curGr++
			right++
		}
		if curGr > days {
			days = curGr
		}
		left++
		curGr--
	}
	return days
}
