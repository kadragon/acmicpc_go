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

	var n, a, b, p int
	fmt.Fscanf(rd, "%d\n", &n)

	d := [101][]int{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d.%d\n", &a, &b)
		d[a] = append(d[a], b)
	}

	for i := 0; i <= 100; i++ {
		sort.Slice(d[i], func(x, y int) bool {
			return d[i][x] < d[i][y]
		})

		for _, v := range d[i] {
			fmt.Fprintf(wr, "%d.%03d\n", i, v)
			p++
			if p == 7 {
				return
			}
		}
	}
}
