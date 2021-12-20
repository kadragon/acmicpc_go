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

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(element interface{}) {
	*h = append(*h, element.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
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

func (h *maxHeap) Push(element interface{}) {
	*h = append(*h, element.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	right := &minHeap{}
	heap.Init(right)

	left := &maxHeap{}
	heap.Init(left)

	var n, t int
	fmt.Fscan(rd, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &t)
		heap.Push(left, t)

		for right.Len() > 0 && (*left)[0] > (*right)[0] {
			p := heap.Pop(left)
			heap.Push(right, p)
		}
		for left.Len() > right.Len()+i%2 {
			p := heap.Pop(left)
			heap.Push(right, p)
		}
		for left.Len()-i%2 < right.Len() {
			p := heap.Pop(right)
			heap.Push(left, p)
		}

		fmt.Fprintf(wr, "%d\n", (*left)[0])
	}
}
