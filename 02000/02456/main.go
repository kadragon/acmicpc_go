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

	var n, a, b, c int
	fmt.Fscan(rd, &n)

	d := make([][5]int, 3)
	d[0][0] = 1
	d[1][0] = 2
	d[2][0] = 3

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a, &b, &c)
		d[0][a]++
		d[0][4] += a
		d[1][b]++
		d[1][4] += b
		d[2][c]++
		d[2][4] += c
	}

	sort.Slice(d, func(i, j int) bool {
		if d[i][4] == d[j][4] {
			if d[i][3] == d[j][3] {
				return d[i][2] > d[j][2]
			}
			return d[i][3] > d[j][3]
		}
		return d[i][4] > d[j][4]
	})

	if d[0][3] == d[1][3] && d[0][2] == d[1][2] {
		fmt.Fprintf(wr, "%d %d\n", 0, d[0][4])
	} else {
		fmt.Fprintf(wr, "%d %d\n", d[0][0], d[0][4])
	}
}
