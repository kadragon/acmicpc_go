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

	d ints
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

type ints [][2]int

func (is ints) Len() int {
	return len(is)
}

func (is ints) Less(i, j int) bool {
	if is[i][0] < is[j][0] {
		return true
	} else if is[i][0] == is[j][0] {
		return is[i][1] < is[j][1]
	} else {
		return false
	}
}

func (is ints) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d = make(ints, n)
	dt := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i][0], &d[i][1])

		if d[i][1] < d[i][0] {
			d[i][0], d[i][1] = d[i][1], d[i][0]
		}

		dt[i] = 1
	}

	sort.Sort(d)

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if d[i][0] >= d[j][0] && d[i][1] >= d[j][1] {
				if dt[i] < dt[j]+1 {
					dt[i] = dt[j] + 1
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if ans < dt[i] {
			ans = dt[i]
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
