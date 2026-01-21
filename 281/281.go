package main

type ZigzagIterator struct {
	cur   int
	size  int
	indxs []int
	vects [][]int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		cur:   0,
		size:  2,
		indxs: []int{0, 0},
		vects: [][]int{v1, v2},
	}
}

func (this *ZigzagIterator) next() int {
	v := this.vects[this.cur]
	i := this.indxs[this.cur]
	res := v[i]
	this.indxs[this.cur]++
	this.cur = (this.cur + 1) % this.size
	return res
}

func (this *ZigzagIterator) hasNext() bool {
	startCur := this.cur
	for this.indxs[this.cur] == len(this.vects[this.cur]) {
		this.cur = (this.cur + 1) % this.size
		if startCur == this.cur {
			return false
		}
	}
	return true
}
