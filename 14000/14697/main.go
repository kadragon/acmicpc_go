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

	var a, b, c, n int
	fmt.Fscan(rd, &a, &b, &c, &n)

	for i := 0; i*a <= n; i++ {
		for j := 0; i*a+j*b <= n; j++ {
			if (n-i*a-j*b)%c == 0 {
				fmt.Fprintf(wr, "%d\n", 1)
				return
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", 0)
}
