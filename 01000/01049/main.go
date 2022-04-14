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

	var n, m, ans int
	fmt.Fscan(rd, &n, &m)

	var a, b int
	ma, mb := 1001, 1001
	for ; m > 0; m-- {
		fmt.Fscan(rd, &a, &b)
		if b*6 < a {
			a = b * 6
		}
		if ma > a {
			ma = a
		}
		if mb > b {
			mb = b
		}
	}

	ans += n / 6 * ma
	if n%6*mb < ma {
		ans += n % 6 * mb
	} else {
		ans += ma
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
