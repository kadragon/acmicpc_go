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

	var x, n, a, b int
	fmt.Fscan(rd, &x, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a, &b)
		x -= a * b
	}

	if x == 0 {
		fmt.Fprintf(wr, "Yes\n")
	} else {
		fmt.Fprintf(wr, "No\n")
	}

}
