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

	var a, b, min, ans int
	fmt.Fscan(rd, &a, &b)

	for i := 1; i*i <= b; i++ {
		if a <= i*i && i*i <= b {
			ans += i * i
			if min == 0 {
				min = i * i
			}
		}
	}

	if min == 0 {
		fmt.Fprintf(wr, "%d\n", -1)
	} else {
		fmt.Fprintf(wr, "%d\n%d\n", ans, min)
	}
}
