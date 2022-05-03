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

	var a, b int
	var p, q, r int

	fmt.Fscanf(rd, "%d:%d:%d\n", &p, &q, &r)
	a = p*3600 + q*60 + r

	fmt.Fscanf(rd, "%d:%d:%d\n", &p, &q, &r)
	b = p*3600 + q*60 + r

	b -= a
	if b < 0 {
		b += 24 * 3600
	}
	if b == 0 {
		b = 24 * 3600
	}

	fmt.Fprintf(wr, "%02d:%02d:%02d\n", b/3600, (b%3600)/60, b%60)
}
