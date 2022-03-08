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

	var n int
	fmt.Fscan(rd, &n)

	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	p := [10][10]int{}
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if d[i][j] == '*' {
				for k := 0; k < 8; k++ {
					nx, ny := i+dx[k], j+dy[k]
					if nx < 0 || ny < 0 || nx == n || ny == n {
						continue
					}
					p[nx][ny]++
				}
			}
		}
	}

	q := make([]string, n)
	mine := false

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &q[i])
	}
	for i := 0; i < n && !mine; i++ {
		for j := 0; j < n && !mine; j++ {
			if q[i][j] == 'x' && d[i][j] == '*' {
				mine = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mine && d[i][j] == '*' {
				fmt.Fprint(wr, "*")
			} else if q[i][j] == 'x' {
				fmt.Fprint(wr, p[i][j])
			} else {
				fmt.Fprintf(wr, "%c", '.')
			}
		}
		fmt.Fprintln(wr)
	}
}
