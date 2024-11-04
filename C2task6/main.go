package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func sumOfProductsOfThree2(numbers []int64) int64 {
	n := len(numbers)
	sum := big.NewInt(0)
	sumN := make([]*big.Int, n-2)
	sumPrefix := big.NewInt(0)
	for i := n - 3; i >= 0; i-- {
		if i == n-3 {
			sum.Add(sum, big.NewInt(0).Mul(big.NewInt(numbers[i+1]), big.NewInt(numbers[i+2])))
			sumPrefix = big.NewInt(numbers[i+2])
			sumN[i] = new(big.Int).Set(sum)
		} else {
			sumPrefix.Add(sumPrefix, big.NewInt(numbers[i+2]))
			sum.Add(sum, big.NewInt(0).Mul(big.NewInt(numbers[i+1]), sumPrefix))
			sumN[i] = new(big.Int).Set(sum)
		}
	}
	sum = big.NewInt(0)
	for i := 0; i < n-2; i++ {
		sum.Add(sum, big.NewInt(0).Mul(big.NewInt(numbers[i]), sumN[i]))
	}
	res := big.NewInt(0).Mod(sum, big.NewInt(1000000007))
	return res.Int64()
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var len int
	fmt.Fscanln(in, &len)
	inputStr := make([]int64, len)
	for i := 0; i < len; i++ {
		fmt.Fscan(in, &inputStr[i])
	}
	result := sumOfProductsOfThree2(inputStr)
	fmt.Fprint(out, result)
}
