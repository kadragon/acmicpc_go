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
	for {
		fmt.Fscan(rd, &n)
		if n == 0 {
			break
		}

		if n > 5000000 {
			n = n / 100 * 80
		} else if n > 1000000 {
			n = n / 100 * 90
		}

		fmt.Fprintf(wr, "%d\n", n)
	}
}
