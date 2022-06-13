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

func main() {
	defer wr.Flush()

	var n, m, x, y, t, c int
	fmt.Fscan(rd, &n, &m)

	x, y, c = 0, 0, 1

	d := [5001][5001]int{}
	p := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for ; d[x][y] == 0 && c < n*m; c++ {
		d[x][y] = c

		nx, ny := x+p[t][0], y+p[t][1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || d[nx][ny] != 0 {
			t = (t + 1) % 4
		}

		x, y = x+p[t][0], y+p[t][1]
	}

	fmt.Fprintf(wr, "%d %d\n", y, x)
}
