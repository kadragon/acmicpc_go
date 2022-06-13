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
	var s string
	var c bool

	for i := 1; ; i++ {
		fmt.Fscan(rd, &a, &s, &b)
		if s == "E" {
			break
		}

		switch s {
		case ">=":
			c = a >= b
		case ">":
			c = a > b
		case "<":
			c = a < b
		case "<=":
			c = a <= b
		case "==":
			c = a == b
		case "!=":
			c = a != b
		}

		fmt.Fprintf(wr, "Case %d: %v\n", i, c)
	}
}
