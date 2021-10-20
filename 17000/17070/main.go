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
	d  [17][17]int
	dt [17][17][3]int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	d = [17][17]int{}
	dt = [17][17][3]int{}
}

func check(i, j, now, k int) {
	if k == 0 {
		if j < n && d[i][j+1] == 0 {
			dt[i][j+1][0] += dt[i][j][now]
		}
	}
	if k == 1 {
		if j < n && i < n && d[i][j+1] == 0 && d[i+1][j] == 0 && d[i+1][j+1] == 0 {
			dt[i+1][j+1][1] += dt[i][j][now]
		}
	}
	if k == 2 {
		if i < n && d[i+1][j] == 0 {
			dt[i+1][j][2] += dt[i][j][now]
		}
	}
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(rd, &d[i][j])
		}
	}
	dt[1][2][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 0; k < 3; k++ {
				if dt[i][j][k] > 0 {
					if k == 0 {
						check(i, j, k, 0)
						check(i, j, k, 1)
					} else if k == 1 {
						check(i, j, k, 0)
						check(i, j, k, 1)
						check(i, j, k, 2)
					} else {
						check(i, j, k, 1)
						check(i, j, k, 2)
					}
				}
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", dt[n][n][0]+dt[n][n][1]+dt[n][n][2])
}
