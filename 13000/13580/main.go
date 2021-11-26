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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var a, b, c int
	fmt.Fscan(rd, &a, &b, &c)

	if a == b || b == c || a == c || a+b == c || b+c == a || a+c == b {
		fmt.Fprintf(wr, "S")
	} else {
		fmt.Fprintf(wr, "N")
	}
}
