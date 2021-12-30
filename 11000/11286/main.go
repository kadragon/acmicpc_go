package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type absHeap []int

func (h absHeap) Len() int {
	return len(h)
}

func (h absHeap) Less(i, j int) bool {
	a, b := math.Abs(float64(h[i])), math.Abs(float64(h[j]))

	if a == b {
		return h[i] < h[j]
	} else {
		return a < b
	}
}

func (h absHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *absHeap) Push(element interface{}) {
	*h = append(*h, element.(int))
}

func (h *absHeap) Pop() interface{} {
	old := *h
	n := len(old)
	element := old[n-1]
	*h = old[0 : n-1]
	return element
}

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, t int
	fmt.Fscan(rd, &n)

	abs := &absHeap{}
	heap.Init(abs)

	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &t)
		if t != 0 {
			heap.Push(abs, t)
		} else {
			if abs.Len() == 0 {
				fmt.Fprintf(wr, "%d\n", 0)
			} else {
				fmt.Fprintf(wr, "%d\n", heap.Pop(abs))
			}
		}
	}
}
