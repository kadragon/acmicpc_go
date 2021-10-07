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

	t := []int{1, 2, 5, 7}
	d := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		d[i] = 100001

		for _, p := range t {
			if i >= p && d[i] > d[i-p] + 1 {
				d[i] = d[i-p] + 1
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", d[n])
}
