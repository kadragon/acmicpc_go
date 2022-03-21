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

	var n int
	var s string

	for {
		fmt.Fscan(rd, &n)
		if n == 0 {
			break
		}

		fmt.Fscan(rd, &s)

		for i := 0; i < n; i++ {
			for j := 0; j < len(s)/n; j++ {
				if j%2 == 0 {
					fmt.Fprintf(wr, "%c", s[j*n+i])
				} else {
					fmt.Fprintf(wr, "%c", s[j*n+(n-i-1)])
				}
			}
		}

		fmt.Fprintln(wr)
	}
}
