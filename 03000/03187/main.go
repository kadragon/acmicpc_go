package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer

	d     []string
	check [251][251]bool
	cnt   [][2]int

	r, c int

	dx, dy []int
)

func f(x, y int) [2]int {

	v, k := 0, 0

	check[x][y] = true

	if d[x][y] == 'v' {
		v++
	} else if d[x][y] == 'k' {
		k++
	}

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if !(nx < 0 || ny < 0 || nx == r || ny == c) {
			if d[nx][ny] != '#' && !check[nx][ny] {
				res := f(nx, ny)
				v += res[0]
				k += res[1]
			}
		}
	}

	return [2]int{v, k}
}

func main() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	fmt.Fscan(rd, &r, &c)

	d = make([]string, r)
	check = [251][251]bool{}
	cnt = make([][2]int, 0)

	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}

	for i := 0; i < r; i++ {
		fmt.Fscan(rd, &d[i])
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if d[i][j] != '#' && !check[i][j] {
				cnt = append(cnt, f(i, j))
			}
		}
	}

	v, k := 0, 0
	for _, b := range cnt {
		if b[1] > b[0] {
			k += b[1]
		} else {
			v += b[0]
		}
	}

	fmt.Fprintf(wr, "%d %d\n", k, v)
}
