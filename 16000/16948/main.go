package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

type pair struct {
	x, y int
}

func main() {
	defer wr.Flush()

	var n, ex, ey, sx, sy int
	fmt.Fscan(rd, &n)

	p := make([]pair, 0)
	dt := [201][201]int{}

	fmt.Fscan(rd, &ex, &ey, &sx, &sy)
	p = append(p, pair{sx, sy})
	dt[sx][sy] = 1

	dx := []int{-2, -2, 0, 0, 2, 2}
	dy := []int{-1, 1, -2, 2, -1, 1}

	for len(p) > 0 {
		x, y := p[0].x, p[0].y
		p = p[1:]

		for i := 0; i < 6; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx < 0 || ny < 0 || nx >= n || ny >= n {
				continue
			}

			if dt[nx][ny] == 0 {
				p = append(p, pair{nx, ny})
				dt[nx][ny] = dt[x][y] + 1

				if nx == ex && ny == ey {
					fmt.Fprintf(wr, "%d\n", dt[nx][ny]-1)
					return
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", -1)
}
