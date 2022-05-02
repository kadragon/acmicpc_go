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

	d := make([]int, 7)
	for i := 0; i < 7; i++ {
		fmt.Fscan(rd, &d[i])
	}

	sort.Ints(d)

	fmt.Fprintf(wr, "%d %d %d\n", d[0], d[1], d[6]-d[0]-d[1])
}
