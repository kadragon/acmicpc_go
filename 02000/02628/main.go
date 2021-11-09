package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var w, h, t, a, b, max int
	fmt.Fscan(rd, &w, &h, &t)

	x := []int{0, w}
	y := []int{0, h}

	for i := 1; i <= t; i++ {
		fmt.Fscan(rd, &a, &b)
		if a == 0 {
			y = append(y, b)
		} else {
			x = append(x, b)
		}
	}

	sort.Ints(x)
	sort.Ints(y)

	for i := 1; i < len(x); i++ {
		for j := 1; j < len(y); j++ {
			p := (x[i] - x[i-1]) * (y[j] - y[j-1])
			if p > max {
				max = p
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
