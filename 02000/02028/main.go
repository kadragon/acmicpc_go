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

	var t int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		fmt.Fprintf(wr, "%s\n", solve())
	}
}

func solve() string {
	var n int
	fmt.Fscan(rd, &n)

	l := 1
	for t := n; t > 0; t /= 10 {
		l *= 10
	}

	if (n*n)%l == n {
		return "YES"
	}
	return "NO"
}
