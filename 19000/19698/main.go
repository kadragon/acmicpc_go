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

	var n, w, h, l int
	fmt.Fscan(rd, &n, &w, &h, &l)

	k := (w / l) * (h / l)
	if k > n {
		k = n
	}

	fmt.Fprintf(wr, "%d\n", k)
}
