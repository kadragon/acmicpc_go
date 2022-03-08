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

	var n, k int
	fmt.Fscan(rd, &n, &k)

	p := (1 + k) * k / 2
	if n < p {
		fmt.Fprintf(wr, "%d\n", -1)
	} else {
		if (n-p)%k == 0 {
			fmt.Fprintf(wr, "%d\n", k-1)
		} else {
			fmt.Fprintf(wr, "%d\n", k)
		}
	}
}
