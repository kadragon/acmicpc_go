package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n, t int
	d    [][]int
	s    string

	dt []int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func f(state, cnt int) int {
	if cnt >= t {
		return 0
	}

	if dt[state] == -1 {
		dt[state] = 987654321

		for i := 0; i < n; i++ {
			if (state & (1 << i)) > 0 {
				for j := 0; j < n; j++ {
					if ((state & (1 << j)) == 0) && i != j {
						dt[state] = min(dt[state], f(state|(1<<j), cnt+1)+d[i][j])
					}
				}
			}
		}
	}

	return dt[state]
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	fmt.Fscan(rd, &n)
	d = make([][]int, n)
	dt = make([]int, 1<<n)

	for i := 0; i < len(dt); i++ {
		dt[i] = -1
	}

	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	fmt.Fscan(rd, &s, &t)
	state, start := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'Y' {
			state |= 1 << i
			start++
		}
	}

	if state == 0 && t > 0 {
		fmt.Fprintf(wr, "%d\n", -1)
	} else {
		fmt.Fprintf(wr, "%d\n", f(state, start))
	}
}
