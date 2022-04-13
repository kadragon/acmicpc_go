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

	d []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var s, c int
	var total int64
	fmt.Fscan(rd, &s, &c)

	d = make([]int, s)
	for i := 0; i < s; i++ {
		fmt.Fscan(rd, &d[i])
		total += int64(d[i])
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})

	t := upper_bound(1, d[s-1], c)

	fmt.Fprintf(wr, "%d\n", total-int64(t*c))
}

func upper_bound(s, e, k int) int {
	var t, m int

	for e > s {
		t = 0
		m = (s + e) / 2
		for _, v := range d {
			t += v / m
		}
		if t >= k {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	return e
}
