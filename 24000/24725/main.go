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

	var n, m, ans int
	fmt.Fscan(rd, &n, &m)

	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	p := [][2]int{{-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}}
	a := " NFP"
	b := " STJ"
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if d[i][j] == 'I' || d[i][j] == 'E' {
				for k := 0; k < 8; k++ {
					for q := 1; q < 4; q++ {
						nx, ny := i+p[k][0]*q, j+p[k][1]*q
						if nx < 0 || ny < 0 || nx == n || ny == m {
							break
						}
						if !(d[nx][ny] == a[q] || d[nx][ny] == b[q]) {
							break
						}
						if q == 3 {
							ans++
						}
					}
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
