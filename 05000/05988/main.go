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

	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &s)
		
		if (s[len(s)-1] - '0') % 2 == 0 {
			fmt.Fprintf(wr, "even\n")
		} else {
			fmt.Fprintf(wr, "odd\n")
		}
	}
}
