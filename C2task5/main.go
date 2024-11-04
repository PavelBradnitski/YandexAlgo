package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var length int
	fmt.Fscanln(in, &length)
	intSlice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(in, &intSlice[i])
	}
	fmt.Fprint(out, Median(intSlice))

}

func Median(numbers []int) string {
	sort.Ints(numbers)
	var sb strings.Builder
	n := len(numbers)
	center := n / 2
	if n%2 == 1 {
		sb.WriteString(fmt.Sprintf("%d ", numbers[center]))
		for offset := 1; offset <= center; offset++ {
			if center-offset >= 0 {
				sb.WriteString(fmt.Sprintf("%d ", numbers[center-offset]))
			}
			if center+offset < n {
				sb.WriteString(fmt.Sprintf("%d ", numbers[center+offset]))
			}
		}
	} else {
		for offset := 0; offset <= center; offset++ {
			if center-1-offset >= 0 {
				sb.WriteString(fmt.Sprintf("%d ", numbers[center-1-offset]))
			}
			if center+offset < len(numbers) {
				sb.WriteString(fmt.Sprintf("%d ", numbers[center+offset]))
			}
		}
	}
	return sb.String()
}
