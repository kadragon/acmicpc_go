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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	a, b := d[0], d[1]
	i, j := 0, n-1

	for i < j {
		if abs(d[i]+d[j]) < abs(a+b) {
			a, b = d[i], d[j]
		}

		if -a == b {
			break
		}

		if -d[i] > d[j] {
			i++
		} else {
			j--
		}
	}

	fmt.Fprintf(wr, "%d %d\n", a, b)
}
