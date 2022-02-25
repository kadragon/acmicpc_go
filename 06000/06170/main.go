package main

import (
	"bufio"
	"fmt"
	"os"
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

	var t int
	fmt.Fscan(rd, &t)

	d := make([][2]int, 46)
	d[1][0] = 1
	d[1][1] = 1

	for i := 2; i <= 45; i++ {
		d[i][0] = d[i-1][0] + d[i-1][1]
		d[i][1] = d[i-1][0]
	}

	var n int
	for i := 1; i <= t; i++ {
		fmt.Fscan(rd, &n)
		fmt.Fprintf(wr, "Scenario #%d:\n%d\n\n", i, d[n][0]+d[n][1])
	}
}
