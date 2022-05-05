package main

import (
	"bufio"
	"fmt"
	"os"
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

	var n, ans int
	fmt.Fscan(rd, &n)

	var t string
	d := make([][26]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		for _, v := range t {
			d[i][int(v-'A')]++
		}
	}

	for i := 1; i < n; i++ {
		m, p := 0, 0

		for j := 0; j < 26; j++ {
			if d[0][j] != d[i][j] {
				t := d[0][j] - d[i][j]
				if t == -1 {
					m++
				} else if t == 1 {
					p++
				} else {
					p = 90
					break
				}
			}
		}

		if !(m > 1 || p > 1) {
			ans++
		}
	}

	fmt.Fprintf(wr, "%d\n", ans)
}
