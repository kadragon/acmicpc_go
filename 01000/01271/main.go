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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var a, b big.Int
	fmt.Fscan(rd, &a, &b)
	fmt.Fprintf(wr, "%d\n%d", new(big.Int).Div(&a, &b), new(big.Int).Mod(&a, &b))
}
