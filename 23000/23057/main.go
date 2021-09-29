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

	var n, t, s int

	d := map[int]bool{0: true}
	l := []int{0}

	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &t)
		s += t
		lenL := len(l)

		for _, p := range l[:lenL] {
			if _, exists := d[p+t]; !exists {
				d[p+t] = true
				l = append(l, p+t)
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", s-len(d)+1)

}
