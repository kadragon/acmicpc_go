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

	var s, m, l int
	fmt.Fscan(rd, &s, &m, &l)

	if s+m*2+l*3 >= 10 {
		fmt.Fprintf(wr, "happy\n")
	} else {
		fmt.Fprintf(wr, "sad\n")
	}
}
