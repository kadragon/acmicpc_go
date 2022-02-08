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

	a, b []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func LowerBound(target int) int {
	low, high, mid := 0, len(b)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if b[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func UpperBound(target int) int {
	low, high, mid := 0, len(b)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if b[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d := [4][4001]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[0][i], &d[1][i], &d[2][i], &d[3][i])
	}

	a = make([]int, n*n)
	b = make([]int, n*n)

	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[k] = d[0][i] + d[1][j]
			b[k] = d[2][i] + d[3][j]
			k++
		}
	}

	sort.Ints(b)

	var ans int64 = 0
	for _, v := range a {
		lower := LowerBound(-v)

		if lower != n*n && -v == b[lower] {
			upper := UpperBound(-v)
			ans += int64(upper - lower)
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
