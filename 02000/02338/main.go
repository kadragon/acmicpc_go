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

	var a, b, c big.Int
	fmt.Fscan(rd, &a, &b)

	c.Add(&a, &b)
	fmt.Fprintln(wr, c.String())

	c.Sub(&a, &b)
	fmt.Fprintln(wr, c.String())

	c.Mul(&a, &b)
	fmt.Fprintln(wr, c.String())
}
