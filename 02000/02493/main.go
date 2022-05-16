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

	d := make([][2]int, 0)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)

		for len(d) > 0 && d[len(d)-1][0] <= t {
			d = d[:len(d)-1]
		}

		if len(d) == 0 {
			fmt.Fprintf(wr, "0 ")
		} else {
			fmt.Fprintf(wr, "%d ", d[len(d)-1][1]+1)
		}

		d = append(d, [2]int{t, i})
	}
}
