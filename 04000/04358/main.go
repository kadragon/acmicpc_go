package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	c := map[string]int{}
	t := []string{}

	for {
		p, _, _ := rd.ReadLine()
		if len(p) == 0 {
			break
		}
		s = string(p)

		if _, ok := c[s]; !ok {
			t = append(t, s)
		}

		c[s]++
		n++
	}

	sort.Strings(t)

	for _, s := range t {
		fmt.Fprintf(wr, "%s %.4f\n", s, float64(c[s])/float64(n)*100)
	}
}
