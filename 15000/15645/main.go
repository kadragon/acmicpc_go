package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd     *bufio.Reader
	wr     *bufio.Writer
	dn, dm [2][3]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	dn = [2][3]int{}
	dm = [2][3]int{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer wr.Flush()

	var n int
	d := make([]int, 3)
	fmt.Fscan(rd, &n)

	x, y := 0, 1

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &d[0], &d[1], &d[2])
		dn[y][0] = min(dn[x][0], dn[x][1]) + d[0]
		dn[y][1] = min(min(dn[x][0], dn[x][1]), dn[x][2]) + d[1]
		dn[y][2] = min(dn[x][1], dn[x][2]) + d[2]

		dm[y][0] = max(dm[x][0], dm[x][1]) + d[0]
		dm[y][1] = max(max(dm[x][0], dm[x][1]), dm[x][2]) + d[1]
		dm[y][2] = max(dm[x][1], dm[x][2]) + d[2]

		x, y = y, x
	}

	a_min, a_max := dn[1][0], dm[1][0]

	for i := 1; i < 3; i++ {
		a_min = min(a_min, dn[1][i])
		a_max = max(a_max, dm[1][i])
	}

	fmt.Fprintf(wr, "%d %d\n", a_max, a_min)
}
