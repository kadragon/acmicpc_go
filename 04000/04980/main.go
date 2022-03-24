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

	p := make([]bool, 1000001)
	p[1] = true

	for i := 2; i < 1000001; i++ {
		if !p[i] {
			for j := 2; i*j < 1000001; j++ {
				p[i*j] = true
			}
		}
	}

	var a, d, n int
	for {
		fmt.Fscan(rd, &a, &d, &n)
		if d == 0 {
			break
		}

		c := 0
		for i := 0; ; i++ {
			if !p[a+d*i] {
				c++
				if c == n {
					fmt.Fprintf(wr, "%d\n", a+d*i)
					break
				}
			}
		}
	}
}
