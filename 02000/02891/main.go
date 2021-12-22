package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	ds [11]int
	dr [11]int

	rec, n int
)

func f(k, r int) {
	if k > n {
		if rec < r {
			rec = r
		}

		return
	}

	if ds[k] == 1 {
		if dr[k-1] == 1 {
			dr[k-1] = 0
			f(k+1, r+1)
			dr[k-1] = 1
		}
		if k+1 <= n && dr[k+1] == 1 {
			dr[k+1] = 0
			f(k+1, r+1)
			dr[k+1] = 1
		}
	}
	f(k+1, r)
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var s, r, t, re int
	fmt.Fscan(rd, &n, &s, &r)

	ds = [11]int{}
	dr = [11]int{}

	for i := 0; i < s; i++ {
		fmt.Fscan(rd, &t)
		ds[t] = 1
	}

	for i := 0; i < r; i++ {
		fmt.Fscan(rd, &t)
		if ds[t] == 1 {
			ds[t] = 0
			re++
		} else {
			dr[t] = 1
		}
	}

	f(1, 0)

	fmt.Fprintf(wr, "%d\n", s-rec-re)
}
