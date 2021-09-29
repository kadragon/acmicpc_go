package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d [21][21]int
	p [4][2]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	d = [21][21]int{}
	p = [4][2]int{{-1, 1}, {0, 1}, {1, 1}, {1, 0}}
}

func c(x, y int) bool {
	if x == 0 || y == 0 || x == 20 || y == 20 {
		return false
	}
	return true
}

func f(x, y int) int {
	var nx, ny int
	var find bool
	six := [2]int{-1, 5}

	for t := 0; t < 4; t++ {
		find = true
		for l := 1; l <= 4; l++ {
			nx = x + l*p[t][0]
			ny = y + l*p[t][1]
			if !c(nx, ny) || d[x][y] != d[nx][ny] {
				find = false
				break
			}
		}
		for l := 0; l < 2 && find; l++ {
			nx = x + six[l]*p[t][0]
			ny = y + six[l]*p[t][1]
			if d[x][y] == d[nx][ny] {
				find = false
				break
			}
		}

		if find {
			return d[x][y]
		}
	}

	return 0
}

func main() {
	defer wr.Flush()

	for i := 1; i <= 19; i++ {
		for j := 1; j <= 19; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	for i := 1; i <= 19; i++ {
		for j := 1; j <= 19; j++ {
			if d[i][j] != 0 {
				r := f(i, j)
				if r != 0 {
					fmt.Fprintf(wr, "%d\n%d %d", r, i, j)
					return
				}
			}
		}
	}

	fmt.Fprintf(wr, "0")
}
