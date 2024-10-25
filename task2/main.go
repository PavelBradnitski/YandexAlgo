package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Println(FindMinComb(a, b, c, d))
}

func FindMinComb(a, b, c, d int) (int, int) {
	var sum, res1, res2 int

	if a == 0 || b == 0 || c == 0 || d == 0 {
		switch {
		case a == 0:
			res1 = 1
			res2 = c + 1
		case b == 0:
			res1 = 1
			res2 = d + 1
		case c == 0:
			res1 = a + 1
			res2 = 1
		case d == 0:
			res1 = b + 1
			res2 = 1
		}
	} else {
		if a > b {
			res1 = a + 1
			res2 = 1
			sum = res1 + res2
		} else {
			res1 = b + 1
			res2 = 1
			sum = res1 + res2
		}

		if c > d {
			if sum > c+2 || sum == 0 {
				res1 = 1
				res2 = c + 1
				sum = res1 + res2
			}
		} else {
			if sum > d+2 || sum == 0 {
				res1 = 1
				res2 = d + 1
				sum = res1 + res2
			}
		}
		if sum > a+c+2 || sum == 0 {
			res1 = a + 1
			res2 = c + 1
			sum = res1 + res2
		}
		if sum > b+d+2 || sum == 0 {
			res1 = b + 1
			res2 = d + 1
		}
	}
	return res1, res2
}
