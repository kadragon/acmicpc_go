package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n, MAX int
	d      [10][10]int
	dt     [1 << 10][10]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	MAX = 10000001
	d = [10][10]int{}
	dt = [1 << 10][10]int{}
}

func TSP(visit, now int) int {
	visit |= (1 << now)

	if visit == (1<<n)-1 {
		if d[now][0] > 0 {
			return d[now][0]
		}
		return MAX
	}

	if dt[visit][now] == 0 {
		dt[visit][now] = MAX

		for i := 0; i < n; i++ {
			if now != i && visit&(1<<i) == 0 && d[now][i] > 0 {
				tmp := TSP(visit, i) + d[now][i]
				if tmp < dt[visit][now] {
					dt[visit][now] = tmp
				}
			}
		}
	}

	return dt[visit][now]
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}

	fmt.Fprintf(wr, "%d\n", TSP(0, 0))
}
