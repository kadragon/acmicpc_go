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
	for {
		var ans int
		var a, b, c float64
		n, _ := fmt.Fscan(rd, &a, &b, &c)
		if n == 0 {
			break
		}

		for a <= c {
			ans++
			a *= (1 + b/100)
		}

		fmt.Fprintf(wr, "%d\n", ans)
	}
}
