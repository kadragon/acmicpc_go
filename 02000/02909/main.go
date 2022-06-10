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

	if k == 0 {
		fmt.Fprintf(wr, "%d\n", n)
		return
	}

	t := k - 1
	for i := 0; i < t; i++ {
		n /= 10
	}

	if n%10 >= 5 {
		n = n/10 + 1
	} else {
		n /= 10
	}

	for i := 0; i <= t; i++ {
		n *= 10
	}

	fmt.Fprintf(wr, "%d\n", n)
}
