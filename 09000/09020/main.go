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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, t int
	fmt.Fscanf(rd, "%d", &t)

	d := [10001]int{}
	d[1] = 2

	for i := 2; i <= 10000; i++ {
		if d[i] == 0 {
			d[i] = 1
			for j := 2; j*i <= 10000; j++ {
				d[i*j] = 2
			}
		}
	}

	for k := 0; k < t; k++ {
		fmt.Fscan(rd, &n)
		for i := n / 2; i > 1; i-- {
			if d[i]+d[n-i] == 2 {
				fmt.Fprintf(wr, "%d %d\n", i, n-i)
				break
			}
		}
	}

}
