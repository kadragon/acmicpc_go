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

	var t, a, b, c, ans int
	var answers []string = []string{"yes", "no"}
	fmt.Fscan(rd, &t)

	for i := 1; i <= t; i++ {
		fmt.Fscan(rd, &a, &b, &c)
		if a > b {
			a, b = b, a
		}
		if b > c {
			b, c = c, b
		}

		if a*a+b*b == c*c && a+b > c {
			ans = 0
		} else {
			ans = 1
		}

		fmt.Fprintf(wr, "Scenario #%d:\n%s\n\n", i, answers[ans])
	}
}
