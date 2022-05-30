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

	var n, t, ans int
	var s string
	fmt.Fscan(rd, &n, &s)

	for _, v := range s {
		if v >= '0' && v <= '9' {
			t *= 10
			t += int(v - '0')
		} else {
			ans += t
			t = 0
		}
	}

	fmt.Fprintf(wr, "%d\n", ans+t)
}
