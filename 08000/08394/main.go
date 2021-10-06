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

	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		a, b = b, (a+b)%10
	}

	fmt.Fprintf(wr, "%d", b)
}
