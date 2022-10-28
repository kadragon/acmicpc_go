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

	var a, b, c, d int64
	fmt.Fscan(rd, &a, &b, &c)

	d = a * b * c

	cnt := [10]int{}

	for d >= 10 {
		cnt[d%10]++
		d /= 10
	}
	cnt[d]++

	for i := 0; i < 10; i++ {
		fmt.Fprintf(wr, "%d\n", cnt[i])
	}
}
