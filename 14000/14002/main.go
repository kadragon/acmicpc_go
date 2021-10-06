package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	dt, data, k []int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func f(p int) {
	if p > 0 {
		f(k[p])
		fmt.Fprintf(wr, "%d ", data[p])
	}
}

func main() {
	defer wr.Flush()

	var n, max int
	fmt.Fscan(rd, &n)

	dt = make([]int, n+1)
	data = make([]int, n+1)
	k = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &data[i])
	}

	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if data[j] < data[i] {
				if dt[i] < dt[j]+1 {
					dt[i] = dt[j] + 1
					k[i] = j

					if dt[max] < dt[i] {
						max = i
						break
					}
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[max])
	f(max)
}
