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

	var n, t, ans int
	fmt.Fscan(rd, &n)

	var plus, minus []int

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if t <= 0 {
			minus = append(minus, t)
		} else {
			plus = append(plus, t)
		}
	}

	sort.Ints(minus)
	sort.Ints(plus)

	for i := 0; i < len(minus); i++ {
		if i+1 < len(minus) {
			ans += minus[i] * minus[i+1]
			i++
		} else {
			ans += minus[i]
		}
	}

	for i := len(plus) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			if plus[i]*plus[i-1] > plus[i]+plus[i-1] {
				ans += plus[i] * plus[i-1]
				i--
			} else {
				ans += plus[i]
			}
		} else {
			ans += plus[i]
		}
	}
	fmt.Fprintf(wr, "%d\n", ans)
}
