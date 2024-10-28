package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x, y int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(DefinePath(x1, y1, x2, y2, x, y))
}

func DefinePath(x1, y1, x2, y2, x, y int) string {
	if x1 <= x && x <= x2 {
		if y < y1 {
			return "S"
		} else if y > y2 {
			return "N"
		}
	}
	if y1 <= y && y <= y2 {
		if x < x1 {
			return "W"
		} else if x > x2 {
			return "E"
		}
	}
	if x < x1 && y < y1 {
		return "SW"
	}
	if x < x1 && y > y2 {
		return "NW"
	}
	if x > x2 && y < y1 {
		return "SE"
	}
	if x > x2 && y > y2 {
		return "NE"
	}
	return ""
}
