package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d []int
)

func binary(a, b, t int) int {
	if a > b {
		return -1
	}
	m := (a + b) / 2
	if d[m] == t {
		return m
	} else if d[m] > t {
		return binary(a, m-1, t)
	}
	return binary(m+1, b, t)
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var t int
	n, m := 1, 1

	for {
		fmt.Fscan(rd, &n, &m)
		if n == 0 && m == 0 {
			break
		}

		cnt := 0
		d = make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &d[i])
		}

		for i := 0; i < m; i++ {
			fmt.Fscan(rd, &t)
			if binary(0, n-1, t) != -1 {
				cnt++
			}
		}

		fmt.Fprintf(wr, "%d\n", cnt)
	}
}
