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

	var a, b big.Int
	fmt.Fscan(rd, &a, &b)

	fmt.Fprintln(wr, a.Add(&a, &b).String())
}
