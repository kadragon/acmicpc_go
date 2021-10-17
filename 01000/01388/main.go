package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	n, m, ans int
	d         [101]string
	check     [101][101]bool
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	d = [101]string{}
	check = [101][101]bool{}
}

func f(x, y int) {
	check[x][y] = true
	if d[x][y] == '-' && y < m-1 {
		if d[x][y] == d[x][y+1] && !check[x][y+1] {
			f(x, y+1)
		}
	} else if d[x][y] == '|' && x < n-1 {
		if d[x+1][y] == '|' && !check[x+1][y] {
			f(x+1, y)
		}
	}
}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n, &m)

	for i := 0; i <= n; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !check[i][j] {
				f(i, j)
				ans++
			}
		}
	}

	fmt.Fprintf(wr, "%#v", ans)
}
