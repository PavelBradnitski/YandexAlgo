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
	fmt.Fprint(out, CountSubarraysWithSum(inputStr, num))

}

func CountSubarraysWithSum(numbers []int, k int) int {
	n := len(numbers)
	prefixSums := make(map[int]int)
	prefixSums[0] = 1 // Инициализируем префиксную сумму 0
	count := 0

	// Вычисляем префиксные суммы и обновляем словарь
	sum := 0
	for i := 0; i < n; i++ {
		sum += numbers[i]
		prefixSums[sum]++ // Увеличиваем счетчик для этой префиксной суммы
	}
	sum = 0
	// Подсчитываем количество подмассивов с суммой k
	for i := 0; i < n; i++ {
		sum += numbers[i]
		count += prefixSums[sum-k] // Ищем префиксную сумму, которая соответствует sum-k
	}

	return count
}

// func SumNum(nums []int, num int) int {
// 	var count, length, sum int
// 	for i := 0; i < len(nums); i++ {
// 		if sum < num {
// 			sum += nums[i]
// 			if sum == num {
// 				if i != 0 {
// 					sum = 0
// 					i -= length
// 					length = 0
// 					count++
// 				} else {
// 					sum = 0
// 					length = 0
// 					count++
// 				}
// 			} else {
// 				if i == len(nums)-1 && sum > num {
// 					sum = 0
// 					i = i - length
// 					length = 0
// 				} else {
// 					length++
// 				}
// 			}
// 		} else if sum == num {
// 			sum = nums[i-length]
// 			i -= length
// 			length = 0
// 			count++
// 		} else {
// 			sum = 0
// 			i = i - length
// 			length = 0
// 		}
// 	}
// 	return count
// }
