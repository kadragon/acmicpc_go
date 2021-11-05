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

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	defer wr.Flush()

	var n, ans int
	fmt.Fscan(rd, &n)

	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}
	sort.Ints(d)

	for i := 0; i < n-1; i++ {
		d[i] = d[i+1] - d[i]
	}

	g := gcd(d[0], d[1])
	for i := 2; i < n-1; i++ {
		if d[i]%g != 0 {
			g = gcd(g, d[i])
		}
	}

	for i := 0; i < n-1; i++ {
		ans += d[i]/g - 1
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
