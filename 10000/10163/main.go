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

	d := [1005][1005]int{}

	var ax, ay, bx, by int

	for k := 1; k <= n; k++ {
		fmt.Fscan(rd, &ax, &bx, &ay, &by)
		for i := 0; i < ay; i++ {
			for j := 0; j < by; j++ {
				d[i+ax][j+bx] = k
			}
		}
	}

	p := make([]int, n+1)

	for i := 0; i < 1005; i++ {
		for j := 0; j < 1005; j++ {
			p[d[i][j]]++
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(wr, "%d\n", p[i])
	}
}
