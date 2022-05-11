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

	var s string
	fmt.Fscan(rd, &s)

	d := [2]int{}
	for _, v := range s {
		d[v-'0']++
	}

	d[0] /= 2
	d[1] /= 2

	for _, v := range s {
		if v == '0' && d[0] > 0 {
			fmt.Fprint(wr, "0")
			d[0]--
		} else if v == '1' {
			if d[1] > 0 {
				d[1]--
			} else {
				fmt.Fprintf(wr, "1")
			}
		}
	}
}
