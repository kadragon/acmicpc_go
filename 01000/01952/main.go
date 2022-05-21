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

	var a, b int
	fmt.Fscan(rd, &a, &b)

	if a > b {
		fmt.Fprintf(wr, "%d\n", b*2-1)
	} else {
		fmt.Fprintf(wr, "%d\n", a*2-2)
	}
}
