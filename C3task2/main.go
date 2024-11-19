package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var length int
	fmt.Fscan(in, &length)
	inputInt := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(in, &inputInt[i])
	}
	fmt.Fprint(out, FindFirstLessNumber(inputInt))
}

type Pair struct {
	Index int
	Value int
}
type Stack []Pair

func (s *Stack) Push(p Pair) {
	*s = append(*s, p)
}

func (s *Stack) Pop() (Pair, bool) {
	if len(*s) == 0 {
		return Pair{}, false
	}
	last := len(*s) - 1
	p := (*s)[last]
	*s = (*s)[:last]
	return p, true
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
func FindFirstLessNumber(input []int) string {
	stack := Stack{}
	sb := strings.Builder{}

	results := make([]int, len(input))
	for i := range results {
		results[i] = -1
	}
	for i := 0; i < len(input); i++ {
		for !stack.Empty() && stack[len(stack)-1].Value > input[i] {
			p, _ := stack.Pop()
			results[p.Index] = i
		}
		stack.Push(Pair{Index: i, Value: input[i]})
	}

	for _, v := range results {
		sb.WriteString(fmt.Sprintf("%d ", v))
	}

	return sb.String()
}
