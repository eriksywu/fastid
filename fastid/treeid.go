package treeid

import (
	"fmt"
	"math"
)

type ArrayTree struct {
	height int
	size int
	data []bool
}

func NewArrayTree(n int) *ArrayTree {
	height := math.Log2(float64(n)) + 1
	size := 2*(n+1) - 1
	data := make([]bool, size)
	for n := range data {
		data[n] = true
	}
	return &ArrayTree{
		data: data,
		size: n,
		height: int(height),
	}
}

func (a ArrayTree) Alloc() (int, error) {
	if !a.data[0] {
		return 0, fmt.Errorf("no more ids available")
	}
	idx := 0
	value := 0
	height := 0
	for  {
		if 2*idx + 1 < len(a.data) && a.data[2*idx + 1]{
			value = value | (0 << (a.height - height - 1))
			idx = 2*idx + 1
		} else if 2*idx + 2 < len(a.data) && a.data[2*idx + 2] {
			value = value | (1 << (a.height - height - 1))
			idx = 2*idx + 2
		} else {
			break
		}
		height++
	}
	a.markParents(idx, false)
	return value, nil
}

func (a *ArrayTree) Free(id int) error {
	if id > a.size {
		return fmt.Errorf("id out of bounds")
	}
	idx := len(a.data) - (a.size + 1 - id)
	a.markParents(idx, true)
	return nil
}

func (a ArrayTree) markParents(idx int, leafValue bool) {
	for idx >= 1 {
		// no children
		if 2*idx + 1 >= len(a.data) {
			a.data[idx] = leafValue
		} else if 2*idx + 2 >= len(a.data) {
			a.data[idx] = a.data[2*idx + 1]
		} else {
			a.data[idx] = a.data[2*idx + 1] || a.data[2*idx + 2]
		}
		idx = (idx - 1)/2
	}
	a.data[0] = a.data[1] || a.data[2]
}


