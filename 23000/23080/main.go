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

	var n int
	var s string
	fmt.Fscan(rd, &n, &s)
	for i, v := range s {
		if i%n == 0 {
			fmt.Fprintf(wr, "%c", v)
		}
	}
}
