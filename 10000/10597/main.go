package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d    []bool
	s    string
	p    []int
	t    int
	find bool
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &s)

	t = len(s)
	if t > 9 {
		t = 9 + (t-9)/2
	}

	d = make([]bool, t+1)
	p = make([]int, t)

	f(s, 0)
}

func f(s string, i int) {
	if len(s) == 0 {
		if !find {
			for _, v := range p {
				fmt.Fprintf(wr, "%d ", v)
			}
			find = true
		}
		return
	}

	a, _ := strconv.Atoi(s[:1])
	if a <= t && !d[a] {
		d[a] = true
		p[i] = a
		f(s[1:], i+1)
		d[a] = false
	}

	if len(s) > 1 {
		b, _ := strconv.Atoi(s[:2])
		if b <= t && !d[b] {
			d[b] = true
			p[i] = b
			f(s[2:], i+1)
			d[b] = false
		}
	}
}
