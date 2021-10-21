package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var n, t int
	var sum int64
	fmt.Fscan(rd, &n)

	d := []int{}
	m := map[int]int{}

	for i := 0; i < n*n; i++ {
		fmt.Fscan(rd, &t)
		sum += int64(t)

		if _, ok := m[t]; !ok {
			m[t] = 1
			d = append(d, t)
		} else {
			m[t]++
		}
	}

	if sum == 0 {
		fmt.Fprintf(wr, "0")
		return
	}

	sort.Ints(d)

	var total int64
	cnt := n * n
	k := 0
	if d[0] == 0 {
		cnt -= m[0]
		k = 1
	}

	for i := 1; i <= d[len(d)-1]; i++ {
		total += int64(cnt)
		if total*2 >= sum {
			fmt.Fprintf(wr, "%d\n", i)
			break
		}

		if d[k] == i {
			cnt -= m[i]
			k++
		}
	}
}
