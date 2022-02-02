package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	checked    [102][102]bool
	d          [102]string
	n, m, w, b int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	checked = [102][102]bool{}
	d = [102]string{}
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func f(x, y int) int {
	var cnt int = 1
	checked[x][y] = true

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || ny < 0 || nx == m || ny == n {
			continue
		}
		if !checked[nx][ny] && d[x][y] == d[nx][ny] {
			cnt += f(nx, ny)
		}
	}

	return cnt
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n, &m)

	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !checked[i][j] {
				cnt := f(i, j)
				if d[i][j] == 'W' {
					w += cnt * cnt
				} else {
					b += cnt * cnt
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d %d\n", w, b)
}
