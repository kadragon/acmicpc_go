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

	var n, c, p, cnt, sam, t int
	fmt.Fscan(rd, &n, &c, &p)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if t > c {
			cnt++
		}
		if t == c {
			sam++
		}
	}

	if cnt+sam >= p {
		fmt.Fprintf(wr, "-1")
	} else {
		fmt.Fprintf(wr, "%d", cnt+1)
	}
}
