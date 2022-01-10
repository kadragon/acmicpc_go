package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	checked [51]bool
	link    [51][51]int

	n, m int
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	checked = [51]bool{}
}

func f(k int) {
	if checked[k] {
		return
	}

	checked[k] = true

	for i := 1; i <= n; i++ {
		if link[k][i] == 1 && !checked[i] {
			f(i)
		}
	}

}

func main() {
	defer wr.Flush()

	fmt.Fscan(rd, &n, &m)

	link = [51][51]int{}

	var k int
	fmt.Fscan(rd, &k)
	sincere := make([]int, k)

	for i := 0; i < k; i++ {
		fmt.Fscan(rd, &sincere[i])
	}

	var q, t int
	party := [51][]int{}
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &q)

		for j := 0; j < q; j++ {
			fmt.Fscan(rd, &t)
			party[i] = append(party[i], t)
			for o := 0; o < j; o++ {
				link[party[i][j]][party[i][o]] = 1
				link[party[i][o]][party[i][j]] = 1
			}
		}
	}

	for i := 0; i < k; i++ {
		f(sincere[i])
	}

	ans := 0
	for i := 0; i < m; i++ {
		for _, v := range party[i] {
			if checked[v] {
				ans--
				break
			}
		}
		ans++
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
