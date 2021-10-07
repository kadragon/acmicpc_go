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

	var t int
	fmt.Fscan(rd, &t)

	for t > 0 {
		t--

		var n, c, a int
		var d []int

		fmt.Fscan(rd, &n)
		d = make([]int, n+1)

		for i := 1; i <= n; i++ {
			fmt.Fscan(rd, &d[i])
		}

		a = 1
		c = 1
		for d[c] != n {
			a++
			c = d[c]

			if a > n {
				a = 0
				break
			}
		}

		fmt.Fprintf(wr, "%d\n", a)
	}
}
