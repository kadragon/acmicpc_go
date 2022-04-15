package main

import (
	"bufio"
	"container/heap"
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

	var n int
	fmt.Fscan(rd, &n)

	d := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i][0], &d[i][1])
	}

	sort.Slice(d, func(i, j int) bool {
		if d[i][0] == d[j][0] {
			return d[i][1] > d[j][1]
		}
		return d[i][0] > d[j][0]
	})

	h := &IntHeap{}
	heap.Init(h)

	ans, k := 0, 0
	for i := d[0][0]; i > 0; i-- {
		for ; k < n; k++ {
			if d[k][0] >= i {
				heap.Push(h, d[k][1])
			} else {
				break
			}
		}

		if h.Len() > 0 {
			ans += heap.Pop(h).(int)
		}
	}

	fmt.Fprintf(wr, "%v\n", ans)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
