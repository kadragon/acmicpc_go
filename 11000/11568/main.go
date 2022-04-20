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

	var n, ans int
	fmt.Fscan(rd, &n)

	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if d[j] < d[i] && p[i] < p[j]+1 {
				p[i] = p[j] + 1
			}
		}

		if p[i] > ans {
			ans = p[i]
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
