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
	d      [16][16]int
	dt     [1 << 16][16]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	d = [16][16]int{}
	dt = [1 << 16][16]int{}

	MAX = 17000000
}

func TSP(visit int64, now int) int {
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
			if i != now && (visit&(1<<i)) == 0 && d[now][i] > 0 {
				tmp := TSP(visit, i) + d[now][i]
				if dt[visit][now] > tmp {
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

// https://withhamit.tistory.com/246
