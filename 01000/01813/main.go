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

	var n, t int
	fmt.Fscan(rd, &n)

	d := make([]int, 51)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		d[t]++
	}

	for i := n; i >= 0; i-- {
		if i == d[i] {
			fmt.Fprintf(wr, "%d\n", i)
			return
		}
	}

	fmt.Fprintf(wr, "%d\n", -1)
}
