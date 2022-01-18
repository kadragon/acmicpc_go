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

func check(n, k int) int {
	c := []int{}
	for n > 0 {
		c = append(c, n%k)
		n /= k
	}

	for i := 0; i < len(c); i++ {
		if c[i] != c[len(c)-i-1] {
			return 0
		}
	}

	return 1
}

func main() {
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(rd, &n)

		var res int = 0
		for i := 2; i <= 64 && res != 1; i++ {
			res = check(n, i)
		}

		fmt.Fprintf(wr, "%d\n", res)
	}
}
