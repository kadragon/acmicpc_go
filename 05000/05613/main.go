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

	var n, t int
	var s string

	fmt.Fscan(rd, &n)

	for {
		fmt.Fscan(rd, &s)
		if s == "=" {
			fmt.Fprintf(wr, "%d\n", n)
			break
		}

		fmt.Fscan(rd, &t)
		switch s {
		case "+":
			n += t
		case "-":
			n -= t
		case "*":
			n *= t
		case "/":
			n /= t
		}
	}
}
