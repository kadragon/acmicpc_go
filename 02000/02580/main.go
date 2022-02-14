package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d          [9][9]int
	cx, cy, cz [9][10]bool
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	cx = [9][10]bool{}
	cy = [9][10]bool{}
	cz = [9][10]bool{}
}

func f(c int) {
	if c == 81 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Fprintf(wr, "%d ", d[i][j])
			}
			fmt.Fprintf(wr, "\n")
		}

		wr.Flush()
		os.Exit(0)
	}

	x, y := c/9, c%9

	if d[x][y] == 0 {
		for i := 1; i <= 9; i++ {
			if !cx[x][i] && !cy[y][i] && !cz[(x/3)*3+y/3][i] {
				cx[x][i] = true
				cy[y][i] = true
				cz[(x/3)*3+y/3][i] = true
				d[x][y] = i

				f(c + 1)

				cx[x][i] = false
				cy[y][i] = false
				cz[(x/3)*3+y/3][i] = false
				d[x][y] = 0
			}
		}
	} else {
		f(c + 1)
	}
}

func main() {
	defer wr.Flush()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(rd, &d[i][j])

			cx[i][d[i][j]] = true
			cy[j][d[i][j]] = true
			cz[(i/3)*3+j/3][d[i][j]] = true
		}
	}

	f(0)
}
