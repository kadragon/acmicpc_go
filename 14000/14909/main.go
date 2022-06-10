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

	var n, cnt int

	for {
		p, _ := fmt.Fscan(rd, &n)
		if p == 0 {
			break
		}
		if n > 0 {
			cnt++
		}
	}

	fmt.Fprintf(wr, "%d\n", cnt)
}
