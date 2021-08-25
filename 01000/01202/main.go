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

type juwel struct {
	m, v int
}

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var n, k int
	fmt.Fscan(rd, &n, &k)

	data := make([]juwel, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &data[i].m, &data[i].v)
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].v > data[j].v {
			return true
		} else {
			return data[i].m < data[j].m
		}
	})

	cost := make([]int, k)

	for i := 0; i < k; i++ {
		fmt.Fscan(rd, &cost[i])
	}

	sort.Ints(cost)

	var ans int64

	for i := 0; i < k; i++ {
		for j := 0; j < len(data); j++ {
			if data[j].v != 0 && data[j].m <= cost[i] {
				ans += int64(data[j].v)
				data[j].v = 0
				break
			}
		}
	}

	wr.WriteString(fmt.Sprintf("%d\n", ans))
}

// TODO: 미해결 문제
