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

	var s string
	var n int
	fmt.Fscan(rd, &s, &n)

	c := make([]string, n)
	d := make([][2]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &c[i], &d[i][0], &d[i][1])
	}

	p := make([][200001]int, 26)
	for i := 0; i < len(s); i++ {
		for j := 0; j < 26; j++ {
			p[j][i+1] = p[j][i]

			if int(s[i]-'a') == j {
				p[j][i+1]++
			}
		}
	}

	for i := 0; i < n; i++ {
		w := int(c[i][0] - 'a')
		fmt.Fprintf(wr, "%d\n", p[w][d[i][1]+1]-p[w][d[i][0]])
	}
}
