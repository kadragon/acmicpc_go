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
		var a, b int
		var s string

		fmt.Fscan(rd, &a, &b, &s)

		for _, v := range s {
			fmt.Fprintf(wr, "%c", (int((v-'A'))*a+b)%26+'A')
		}
		fmt.Fprintln(wr)
	}
}
