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
	var p, q, r int
	fmt.Fscan(rd, &a, &b, &c, &n)

	max := 0

	for i := 1; i <= n; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			fmt.Fscan(rd, &p, &q, &r)
			sum += p*a + q*b + r*c
		}
		if max < sum {
			max = sum
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}
