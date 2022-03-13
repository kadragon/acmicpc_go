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

	for i := 0; i < t; i++ {
		solve()
		fmt.Fprintln(wr)
	}
}

func solve() {
	var n int
	var c string

	fmt.Fscan(rd, &n, &c)
	switch c {
	case "C":
		var d string
		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &d)
			fmt.Fprintf(wr, "%d ", d[0]-'A'+1)
		}
	case "N":
		var d int
		for i := 0; i < n; i++ {
			fmt.Fscan(rd, &d)
			fmt.Fprintf(wr, "%c ", d+'A'-1)
		}
	}
}
