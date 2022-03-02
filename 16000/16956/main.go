package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var r, c int
	fmt.Fscan(rd, &r, &c)

	d := make([]string, r)
	for i := 0; i < r; i++ {
		fmt.Fscan(rd, &d[i])
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if d[i][j] == 'S' {
				for p := 0; p < 4; p++ {
					nx, ny := i+dx[p], j+dy[p]
					if nx < 0 || nx == r || ny < 0 || ny == c {
						continue
					}
					if d[nx][ny] == 'W' {
						fmt.Fprintf(wr, "0")
						return
					}
				}
			}
		}
		d[i] = strings.ReplaceAll(d[i], ".", "D")
	}

	fmt.Fprintf(wr, "1\n")
	for i := 0; i < r; i++ {
		fmt.Fprintf(wr, "%s\n", d[i])
	}
}
