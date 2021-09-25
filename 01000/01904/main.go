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

	var n, c int
	a := 0
	b := 1

	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		c = a + b
		if c >= 15746 {
			c -= 15746
		}
		a = b
		b = c
	}

	fmt.Fprintf(wr, "%d", c)
}
