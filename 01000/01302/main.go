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
	fmt.Fscan(rd, &n)

	d := map[string]int{}
	books := make([]string, 0)

	var t string
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if _, ok := d[t]; !ok {
			books = append(books, t)
		}

		d[t]++
	}

	sort.Strings(books)

	var max int
	for _, v := range books {
		if d[v] > max {
			max = d[v]
			t = v
		}
	}

	fmt.Fprintf(wr, "%s\n", t)
}
