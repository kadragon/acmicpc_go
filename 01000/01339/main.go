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

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	var n int
	var s string
	fmt.Fscan(rd, &n)

	a := make([]int, 26)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &s)
		k := 1
		for j := len(s) - 1; j >= 0; j-- {
			a[s[j]-'A'] += k
			k *= 10
		}
	}

	sort.Ints(a)

	ans, k := 0, 9
	for i := 25; i > 16; i-- {
		ans += a[i] * k
		k--
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
