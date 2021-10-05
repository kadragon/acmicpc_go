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

	var a, b string
	p := []int{3, 2, 1, 2, 3, 3, 2, 3, 3, 2, 2, 1, 2, 2, 1, 2, 2, 2, 1, 2, 1, 1, 1, 2, 2, 1}
	fmt.Fscan(rd, &a, &b)

	d := make([]int, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		d[i*2] = p[a[i]-'A']
		d[i*2+1] = p[b[i]-'A']
	}

	for i := 0; i < len(a)*2-2; i++ {
		for j := 0; j < len(a)*2-1-i; j++ {
			d[j] = (d[j] + d[j+1]) % 10
		}
	}

	fmt.Fprintf(wr, "%d%d\n", d[0], d[1])
}
