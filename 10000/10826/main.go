package main

import (
	"bufio"
	"fmt"
	"math/big"
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

	a, b := big.NewInt(0), big.NewInt(1)

	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	fmt.Fprintln(wr, a)
}
