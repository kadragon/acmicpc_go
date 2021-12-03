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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var n, k int
	fmt.Fscan(rd, &n, &k)

	d := make([]int, n+1)
	p := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
		d[i] += d[i-1]
	}

	if d[n] == 0 {
		fmt.Fprintf(wr, "SAD")
		return
	}

	max, cnt := 0, 0
	for i := k; i <= n; i++ {
		p[i] = d[i] - d[i-k]
		if p[i] > max {
			max = p[i]
			cnt = 1
		} else if p[i] == max {
			cnt++
		}
	}

	fmt.Fprintf(wr, "%d\n%d", max, cnt)
}
