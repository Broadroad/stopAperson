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
)

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

// getNumberOfPoint return the number of the points
func getNumberOfPoint(startX, startY int) int {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			tx := startX + dx
			ty := startY + dy

			str := string(tx) + "_" + string(ty)
			if isLessThan21(tx, ty) && flags[str] == false {
				flags[str] = true
				ans++
				getNumberOfPoint(tx, ty)
			}
		}
	}
	return ans
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
