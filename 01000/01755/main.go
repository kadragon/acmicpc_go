package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

type pair struct {
	a int
	s string
}

func main() {
	defer wr.Flush()

	alpha := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var m, n int
	fmt.Fscan(rd, &m, &n)

	d := []pair{}
	for i := m; i <= n; i++ {
		t := i
		p := ""
		for t > 0 {
			p = alpha[t%10] + " " + p
			t /= 10
		}
		d = append(d, pair{i, strings.TrimRight(p, " ")})
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i].s < d[j].s
	})

	for i, v := range d {
		fmt.Fprintf(wr, "%d ", v.a)
		if (i+1)%10 == 0 {
			fmt.Fprintln(wr)
		}
	}
}
