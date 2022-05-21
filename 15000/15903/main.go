package main

import (
	"bufio"
	"container/heap"
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

	var n, m, t int
	fmt.Fscan(rd, &n, &m)

	h := &minheap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		heap.Push(h, t)
	}

	for i := 0; i < m; i++ {
		t = heap.Pop(h).(int) + heap.Pop(h).(int)
		heap.Push(h, t)
		heap.Push(h, t)
	}

	ans := 0
	for _, v := range *h {
		ans += v
	}

	fmt.Fprintf(wr, "%d\n", ans)
}

type minheap []int

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
