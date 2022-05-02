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

	var n, max int
	fmt.Fscan(rd, &n)

	d := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &d[i][0], &d[i][1])
	}

	sort.Slice(d, func(i, j int) bool {
		if d[i][0] == d[j][0] {
			return d[i][1] < d[j][1]
		}

		return d[i][0] < d[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	for _, v := range d {
		heap.Push(h, v[1])

		for h.Len() > 0 {
			t := (*h)[0]
			if v[0] >= t {
				heap.Pop(h)
			} else {
				break
			}
		}

		if h.Len() > max {
			max = h.Len()
		}
	}

	fmt.Fprintf(wr, "%d\n", max)
}

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
