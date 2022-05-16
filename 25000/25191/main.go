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

	var n, a, b int
	fmt.Fscan(rd, &n, &a, &b)

	b += a / 2

	if b > n {
		b = n
	}

	fmt.Fprintf(wr, "%d\n", b)
}
