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

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	fmt.Fscanf(rd, "%d %d", &n, &m)

	g := gcd(n, m)

	for i := 1; i <= g; i++ {
		if g%i == 0 && n%i == 0 && m%i == 0 {
			fmt.Fprintf(wr, "%d %d %d\n", i, n/i, m/i)
		}
	}
}
