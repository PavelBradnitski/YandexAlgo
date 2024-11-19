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
	var length, num int
	fmt.Fscanln(in, &length, &num)
	var inpStr string
	fmt.Fscanln(in, &inpStr)
	fmt.Fprint(out, Censor(inpStr, num))

}

func Censor(input string, n int) int {
	if len(input) == 1 {
		return 1
	}
	var aCount, curGr, left, right int
	var outlen, maxlen, bCount int
	for right < len(input) {
		if curGr > n {
			if input[left] == 'a' {
				if curGr >= bCount {
					curGr -= bCount
				} else {
					curGr = 0
				}
				aCount--
				if aCount == 0 {
					bCount = 0
				}
			} else if input[left] == 'b' {
				bCount--
			}
			left++
			if left == len(input)-1 {
				break
			}
			outlen--
			continue
		}
		if input[right] == 'a' {
			aCount++
			right++
			outlen++
			if outlen > maxlen && curGr <= n {
				maxlen = outlen
			}
			continue
		} else if input[right] == 'b' {
			curGr += aCount
			if aCount != 0 {
				bCount++
			}
			if curGr > n {
				if input[left] == 'a' {
					if curGr > bCount {
						curGr -= bCount
					} else {
						curGr = 0
					}
					aCount--
					if aCount == 0 {
						bCount = 0
					}
				} else if input[left] == 'b' {
					bCount--
				}

				left++
				if left >= len(input)-1 {
					break
				}
				curGr -= aCount + 1
				outlen -= 2
				continue
			}
			outlen++
			if outlen > maxlen && curGr <= n {
				maxlen = outlen
			}
			if right != len(input)-1 || left == len(input)-1 {
				right++
			} else {
				left++
				outlen--
			}
			continue
		} else {
			outlen++
			if outlen > maxlen && curGr <= n {
				maxlen = outlen
			}
			if right != len(input)-1 || left == len(input)-1 {
				right++
			} else {
				left++
				outlen--
			}

		}
	}
	return maxlen
}
