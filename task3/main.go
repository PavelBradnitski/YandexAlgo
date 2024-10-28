package main

import "fmt"

func main() {
	var len int

	fmt.Scanln(&len)
	inputStr := make([]string, len)
	for i := 0; i < len; i++ {
		fmt.Scanln(&inputStr[i])
	}
	fmt.Println(inputStr)
	res := CheckO(inputStr)
	fmt.Println(res)
	//fmt.Println(DefinePath(x1, y1, x2, y2, x, y))
}

func CheckI(strSlice []string) bool {
	// находим x1,y1,x2,y2
	x1, y1, x2, y2 := FindBorder(strSlice, rune('#'))
	// проверяем чтобы в промежутке x1-x2, y1-y2 были значения только '#', а в остальных местах '.'
	if checkOutside(strSlice, x1, y1, x2, y2) {
		for x, v := range strSlice {
			if x < x1 || x > x2 {
				continue
			}
			for y, v1 := range v {
				if y < y1 || y > y2 {
					continue
				}
				if v1 != rune('#') {
					return false
				}
			}
		}
	} else {
		return false
	}
	return true
}
func CheckO(strSlice []string) bool {
	// находим x1,y1,x2,y2
	x1, y1, x2, y2 := FindBorder(strSlice, rune('#'))
	if x1 >= x2 || y1 >= y2 {
		return false
	}
	if checkOutside(strSlice, x1, y1, x2, y2) {
		newStrSlice := make([]string, len(strSlice))
		for i := x1; i <= x2; i++ {
			newStrSlice[i] = strSlice[i][y1 : y2+1]
		}
		x1New, y1New, x2New, y2New := FindBorder(newStrSlice, rune('.'))
		if x1New > x2New || y1New > y2New {
			return false
		}
		if checkBorder(newStrSlice) {
			return checkInsideForCandO(newStrSlice, x1New, y1New, x2New, y2New)
		} else {
			return false
		}

	}
	return false
}

func CheckC(strSlice []string) bool {
	// находим x1,y1,x2,y2
	x1, y1, x2, y2 := FindBorder(strSlice, rune('#'))
	if x1 >= x2 || y1 >= y2 {
		return false
	}
	if checkOutside(strSlice, x1, y1, x2, y2) {
		newStrSlice := make([]string, len(strSlice))
		for i := x1; i <= x2; i++ {
			newStrSlice[i] = strSlice[i][y1 : y2+1]
		}
		x1New, y1New, x2New, y2New := FindBorder(newStrSlice, rune('.'))
		if x1New > x2New || y1New > y2New {
			return false
		}
		if checkBorderForC(newStrSlice) {
			return checkInsideForCandO(newStrSlice, x1New, y1New, x2New, y2New)
		} else {
			return false
		}

	}
	return false
}
func FindBorder(strSlice []string, r rune) (x1, y1, x2, y2 int) {
	// находим x1,y1
L:
	for x, v := range strSlice {
		for y, v1 := range v {
			if v1 == r {
				x1 = x
				y1 = y
				break L
			}
		}
	}
	// находим x2, y2
	for x, v := range strSlice {
		if x <= x1 {
			continue
		}
		for y, v1 := range v {
			if y <= y1 {
				continue
			}
			if v1 == r {
				x2 = x
				y2 = y
			} else {
				break
			}
		}
	}
	return x1, y1, x2, y2
}

func checkOutside(strSlice []string, x1, y1, x2, y2 int) bool {
	for x, v := range strSlice {
		for y, v1 := range v {
			if x < x1 || x > x2 || y < y1 || y > y2 {
				if v1 != rune('.') {
					return false
				}
			}
		}

	}
	return true
}
func checkInsideForCandO(strSlice []string, x1, y1, x2, y2 int) bool {
	for x, v := range strSlice {
		if x == 0 || x == len(strSlice) {
			for _, v1 := range v {
				if v1 != rune('#') {
					return false
				}
			}
		}
		if x < x1 || x > x2 {
			continue
		}
		for y, v1 := range v {
			if (y < y1 || y > y2) && v1 == rune('#') {
				continue
			} else if (y < y1 || y > y2) && v1 == rune('.') || v1 != rune('.') {
				return false
			}
		}
	}
	return true
}

// проверяет границу прямоугольника, предполагая что начало рамки с (0,0)
func checkBorder(strSlice []string) bool {
	for x, v := range strSlice {
		if x == 0 || x == len(strSlice)-1 {
			for _, v1 := range v {
				if v1 != rune('#') {
					return false
				}
			}
			continue
		}
		for y, v1 := range v {
			if (y == 0 || y == len(strSlice[x])-1) && v1 != rune('#') {
				return false
			}
		}
	}
	return true
}

// проверяет границу прямоугольника, предполагая что начало рамки с (0,0)
func checkBorderForC(strSlice []string) bool {
	for x, v := range strSlice {
		if x == 0 || x == len(strSlice)-1 {
			for y, v1 := range v {
				if y != 0 && v1 != rune('#') || (x == len(strSlice)-1 && y == 0 && v1 != rune('#')) {
					return false
				}
			}
			continue
		}
		for y, v1 := range v {
			if y == 0 && v1 != rune('#') {
				return false
			}
		}
	}
	return true
}
