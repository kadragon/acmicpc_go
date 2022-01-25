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

	var k, l int
	fmt.Fscan(rd, &k, &l)

	d := map[string]int{}
	t := [500001]string{}

	var s string
	for i := 1; i <= l; i++ {
		fmt.Fscan(rd, &s)
		if _, ok := d[s]; ok {
			t[d[s]] = ""
		}

		d[s], t[i] = i, s
	}

	for i := 1; i <= l && k > 0; i++ {
		if t[i] != "" {
			fmt.Fprintf(wr, "%s\n", t[i])
			k--
		}
	}
}
