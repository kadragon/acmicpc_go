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

	var t int
	fmt.Fscan(rd, &t)

	for p := 0; p < t; p++ {
		var m, sum int64
		var d = [6]int{}

		fmt.Fscan(rd, &m)
		for i := 0; i < 6; i++ {
			fmt.Fscan(rd, &d[i])
			sum += int64(d[i])
		}

		for i := 1; ; i++ {
			if sum > m {
				fmt.Fprintf(wr, "%d\n", i)
				break
			}
			sum *= 4
		}
	}
}
