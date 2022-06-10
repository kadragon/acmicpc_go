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

	var n, a int
	var s string
	fmt.Fscan(rd, &n, &s)
	for _, v := range s {
		if v != 'C' {
			a++
			n--
		}
	}

	ans := n / (a + 1)
	if n%(a+1) != 0 {
		ans++
	}
	fmt.Fprintf(wr, "%d\n", ans)
}
