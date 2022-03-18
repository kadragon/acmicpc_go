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
	fmt.Fscan(rd, &n)

	d := make([]int, 1<<n-1)
	for i := 0; i < 1<<n-1; i++ {
		fmt.Fscan(rd, &d[i])
	}

	p := [][]int{}
	p = append(p, d)
	a := []int{}

	for len(p) > 0 {
		now := p[0]
		l := len(now)
		a = append(a, now[l/2])
		if l > 1 {
			p = append(p, now[:l/2])
			p = append(p, now[l/2+1:])
		}
		p = p[1:]
	}

	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<i; j++ {
			fmt.Fprintf(wr, "%d ", a[c])
			c++
		}
		fmt.Fprintln(wr)
	}
}
