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

	var n, ans int
	fmt.Fscan(rd, &n)

	d := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i][0], &d[i][1])
	}

	sort.Slice(d, func(i, j int) bool {
		if d[i][0] != d[j][0] {
			return d[i][0] > d[j][0]
		} else {
			return d[i][1] < d[j][1]
		}
	})

	for i := 5; i < n; i++ {
		if d[4][0] == d[i][0] {
			ans++
		} else {
			break
		}
	}

	fmt.Fprintf(wr, "%v\n", ans)
}
