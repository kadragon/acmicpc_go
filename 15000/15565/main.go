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

	var n, k, t int
	fmt.Fscan(rd, &n, &k)

	c := make([]int, 0)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if t == 1 {
			c = append(c, i)
		}
	}

	var ans int = 1000001
	if len(c) < k {
		ans = -1
	} else {
		for i := k - 1; i < len(c); i++ {
			t = c[i] - c[i-k+1] + 1
			if ans > t {
				ans = t
			}
		}
	}

	fmt.Fprintf(wr, "%v\n", ans)
}
