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

type t struct {
	cost int
	time int
}

func main() {
	defer wr.Flush()

	var n, a, b, c, ans int
	fmt.Fscan(rd, &n)

	d := make([]t, 0)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a)
		if a == 1 {
			fmt.Fscan(rd, &b, &c)
			d = append(d, t{b, c})
		}

		p := len(d) - 1
		if p >= 0 {
			d[p].time--
			if d[p].time == 0 {
				ans += d[p].cost
				d = d[:p]
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
