package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n  int
	d  []string
	dt [][16]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func f(state, now int) int {
	if state == (1<<n)-1 {
		return 0
	}

	if dt[state][now] == 0 {
		for i := 0; i < n; i++ {
			if (state&(1<<i)) == 0 && d[now][len(d[now])-1] == d[i][0] {
				dt[state][now] = max(dt[state][now], f(state|(1<<i), i)+len(d[i]))
			}
		}
	}

	return dt[state][now]
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	fmt.Fscan(rd, &n)
	d = make([]string, n)
	dt = make([][16]int, 1<<n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		tmp := f(1<<i, i) + len(d[i])
		if ans < tmp {
			ans = tmp
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
