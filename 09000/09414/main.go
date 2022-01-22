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

func pow(a, b int) int64 {
	var t int64 = 1
	for i := 0; i < b; i++ {
		t *= int64(a)
	}

	return t
}

func solve() {
	d := []int{}
	var n int
	for {
		fmt.Fscan(rd, &n)
		if n == 0 {
			break
		}
		d = append(d, n)
	}

	sort.Ints(d)

	var ans int64 = 0
	l := len(d)
	for i, v := range d {
		ans += 2 * pow(v, l-i)
		if ans > 5000000 {
			fmt.Fprintf(wr, "Too expensive\n")
			return
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}

func main() {
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		solve()
	}
}
