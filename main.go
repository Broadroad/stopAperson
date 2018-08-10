package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	startx int
	starty int
	h      bool
	ans    = 0
	flags  map[string]bool // x_y means (x,y) -> bool (already or not)
	max    = 777           // (-777, -777), (-777, 0) (777, 0) (777, 777) is the biggest area
	min    = -777
)

type Position struct {
	X int
	Y int
}

func init() {
	flag.IntVar(&startx, "x", 0, "startx is start x")
	flag.IntVar(&starty, "y", 0, "starty is start y")
}

// isGreatThan21 return true if (x, y) less than 21
func isLessThan21(x, y int) bool {
	if x < 0 {
		x = 0 - x
	}

	if y < 0 {
		y = 0 - y
	}
	xstr := strconv.Itoa(x)
	ystr := strconv.Itoa(y)

	sum := 0
	for _, sx := range xstr {
		sum += int(sx - '0')
	}
	for _, sy := range ystr {
		sum += int(sy - '0')
	}

	if sum <= 21 {
		return true
	}
	return false
}

func getNumberOfPoint(startX, startY int) int {
	getNumberOfPointRecursive(startX, startY)
	return ans
}

// getNumberOfPointRecursive return the number of the points
// Recursive resolution
func getNumberOfPointRecursive(startX, startY int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			tx := startX + dx
			ty := startY + dy

			// rune, string
			stx := strconv.Itoa(tx)
			sty := strconv.Itoa(ty)
			str := stx + "_" + sty

			if tx >= min && tx <= max && ty >= min && ty <= max && isLessThan21(tx, ty) && flags[str] == false {
				flags[str] = true
				ans++
				getNumberOfPoint(tx, ty)
			}
			flags[str] = true
		}
	}
}

// getNumberOfPointNoRecursive
func getNumberOfPointNoRecursive(startX, startY int) int {
	// TODO NoRecursive
	return 1
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	flags = make(map[string]bool)
	ret := getNumberOfPoint(startx, starty)
	fmt.Println(ret)
}
