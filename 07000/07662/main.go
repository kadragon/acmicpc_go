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
	defer wr.Flush()

	var t int
	fmt.Fscan(rd, &t)

	for ; t > 0; t-- {
		solve()
	}
}

func solve() {
	var n int
	fmt.Fscan(rd, &n)

	minheap := &minHeap{}
	maxheap := &maxHeap{}

	heap.Init(minheap)
	heap.Init(maxheap)

	inData := map[int]int{}

	var s string
	var p int
	for ; n > 0; n-- {
		fmt.Fscan(rd, &s, &p)
		if s == "I" {
			heap.Push(minheap, p)
			heap.Push(maxheap, p)
			inData[p]++
		} else {
			if p < 0 {
				for minheap.Len() > 0 {
					tmp := heap.Pop(minheap).(int)
					if inData[tmp] > 0 {
						inData[tmp]--
						break
					}
				}
			} else {
				for maxheap.Len() > 0 {
					tmp := heap.Pop(maxheap).(int)
					if inData[tmp] > 0 {
						inData[tmp]--
						break
					}
				}
			}
		}
	}

	if maxheap.Len() == 0 || minheap.Len() == 0 {
		fmt.Fprintf(wr, "EMPTY\n")
		return
	}

	for maxheap.Len() > 0 {
		tmp := heap.Pop(maxheap).(int)
		if inData[tmp] == 0 {
			if maxheap.Len() == 0 {
				fmt.Fprintf(wr, "EMPTY\n")
				return
			}

			continue
		}

		fmt.Fprintf(wr, "%d ", tmp)
		break
	}

	for {
		tmp := heap.Pop(minheap).(int)
		if inData[tmp] == 0 {
			continue
		}

		fmt.Fprintf(wr, "%d\n", tmp)
		break
	}
}
