package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	n := nextInt()

	d := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		d[i*i] = 1
	}

	for i := 2; i <= n; i++ {
		if d[i] > 0 {
			continue
		}

		d[i] = 5
		for j := int(math.Sqrt(float64(i))); j > 0; j-- {
			d[i] = min(d[i], d[j*j]+d[i-j*j])
			if d[i] == 2 {
				break
			}
		}
	}

	fmt.Fprintf(wr, "%d\n", d[n])
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())

	return v
}
