package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d  [102][102]int
	dt [102][102]int

	n, m, direct int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func f(x, y int) {
	if x == n && y == m {
		return
	}

	if d[x][y] == 2 {
		if dt[n][m] == 0 || dt[n][m] > dt[x][y]+direct {
			dt[n][m] = dt[x][y] + direct
		}
		return
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]

		if nx == 0 || ny == 0 || nx > n || ny > m {
			continue
		}

		if d[nx][ny] != 1 {
			if dt[nx][ny] == 0 || dt[nx][ny] > dt[x][y]+1 {
				dt[nx][ny] = dt[x][y] + 1
				f(nx, ny)
			}
		}
	}
}

func main() {
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &n, &m, &t)

	d = [102][102]int{}
	dt = [102][102]int{}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(rd, &d[i][j])
			if d[i][j] == 2 {
				direct = n - i + m - j
			}
		}
	}

	dt[1][1] = 1
	f(1, 1)

	if dt[n][m]-1 <= t && dt[n][m] > 0 {
		fmt.Fprintf(wr, "%d\n", dt[n][m]-1)
	} else {
		fmt.Fprintf(wr, "Fail\n")
	}
}
