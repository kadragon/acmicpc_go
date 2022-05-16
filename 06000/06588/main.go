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

	d := [1000001]bool{true, true}
	for i := 2; i < 1000001; i++ {
		if !d[i] {
			for j := 2; i*j <= 1000000; j++ {
				d[i*j] = true
			}
		}
	}

	var n int

	for {
		fmt.Fscan(rd, &n)
		if n == 0 {
			break
		}

		for i := 2; i < 1000001; i++ {
			if !d[i] && !d[n-i] {
				fmt.Fprintf(wr, "%d = %d + %d\n", n, i, n-i)
				break
			}
		}
	}
}
