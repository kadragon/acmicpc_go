package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n, ans int
	d, dt  [501][501]int
	dx, dy []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	d = [501][501]int{}
	dt = [501][501]int{}

	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func f(x, y int) int {
	if dt[x][y] == 0 {
		dt[x][y] = 1

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]

			if nx > 0 && ny > 0 && nx <= n && ny <= n {
				if d[x][y] > d[nx][ny] {
					dt[x][y] = max(dt[x][y], f(nx, ny)+1)
				}
			}
		}
	}

	return dt[x][y]
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			res := f(i, j)
			if ans < res {
				ans = res
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
