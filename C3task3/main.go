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
	var length, dist int
	fmt.Fscan(in, &length, &dist)
	intSlice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(in, &intSlice[i])
	}
	output := MinOnDist(intSlice, dist)
	for _, v := range output {
		fmt.Fprintln(out, v)
	}
}

func MinOnDist(arr []int, k int) []int {
	var res []int
	var deque []int

	for i := 0; i < len(arr); i++ {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		for len(deque) > 0 && arr[deque[len(deque)-1]] >= arr[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			res = append(res, arr[deque[0]])
		}
	}

	return res
}
