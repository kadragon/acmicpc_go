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

	var n int
	fmt.Fscan(rd, &n)

	d := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	var ans, div int64
	fmt.Fscan(rd, &div)

	for _, v := range d {
		k := v / div
		if v%div > 0 {
			k++
		}
		ans += k * div
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
