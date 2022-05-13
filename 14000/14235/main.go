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

	var n, t, p int
	fmt.Fscan(rd, &n)

	h := &maxHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)

		if t == 0 {
			if h.Len() > 0 {
				fmt.Fprintf(wr, "%d\n", heap.Pop(h).(int))
			} else {
				fmt.Fprintf(wr, "-1\n")
			}
		} else {
			for j := 0; j < t; j++ {
				fmt.Fscan(rd, &p)
				heap.Push(h, p)
			}
		}
	}
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
