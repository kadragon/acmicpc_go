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

	var sh, sm, eh, em, qh, qm int
	fmt.Fscanf(rd, "%d:%d %d:%d %d:%d\n", &sh, &sm, &eh, &em, &qh, &qm)

	s := sh*60 + sm
	e := eh*60 + em
	q := qh*60 + qm

	d := map[string]bool{}

	var h, m, ans int
	var nick string

	for {
		if n, _ := fmt.Fscanf(rd, "%d:%d %s\n", &h, &m, &nick); n == 0 {
			break
		}

		c := h*60 + m

		if c <= s {
			d[nick] = true
		} else if c >= e && c <= q {
			if d[nick] {
				d[nick] = false
				ans++
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
