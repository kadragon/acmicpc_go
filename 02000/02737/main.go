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
		var n, p int
		fmt.Fscan(rd, &n)

		for i := 1; n > 0; i++ {
			n -= i
			if n%i == 0 {
				p++
			}
		}

		fmt.Fprintf(wr, "%d\n", p-1)
	}
}
