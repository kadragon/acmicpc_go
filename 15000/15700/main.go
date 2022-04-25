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

	var a, b, ans int64
	fmt.Fscan(rd, &a, &b)

	if a*b%2 == 0 {
		ans = a * b / 2
	} else {
		ans = (a*b - 1) / 2
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
