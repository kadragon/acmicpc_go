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
	d := map[string]string{}
	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &s)

		t := []byte(s)
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})

		p := string(t)
		if d[p] == "" {
			d[p] = s
			fmt.Fprintf(wr, "%s\n", s)
		}
	}
}
