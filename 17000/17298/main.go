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

	d := [1000001]int{}
	a := [1000001]int{}
	q := [1000001]int{}
	key := -1

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := n - 1; i >= 0; i-- {
		for key >= 0 {
			if d[i] >= q[key] {
				key--
			} else {
				break
			}
		}
		if key >= 0 {
			a[i] = q[key]
		} else {
			a[i] = -1
		}

		key++
		q[key] = d[i]
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(wr, "%d ", a[i])
	}
}
