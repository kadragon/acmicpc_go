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

	var n, max int
	fmt.Fscan(rd, &n)

	d := make([]int, n)
	p := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
		p[i] = 1
		for j := 0; j < i; j++ {
			if d[i] > d[j] && p[i] < p[j]+1 {
				p[i] = p[j] + 1
			}
		}

		if p[i] > max {
			max = p[i]
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
