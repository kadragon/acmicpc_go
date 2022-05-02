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

	d := make([]int, 10000)
	d[0] = 1
	for i := 1; i <= 9999; i++ {
		d[i] = i * d[i-1]
		for d[i]%10 == 0 {
			d[i] /= 10
		}

		d[i] %= 100000
	}

	var n int
	for {
		l, _ := fmt.Fscan(rd, &n)
		if l == 0 {
			break
		}

		fmt.Fprintf(wr, "%5d -> %d\n", n, d[n]%10)
	}
}
