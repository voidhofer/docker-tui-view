package main

import (
	"sync"
)

type randomItemGenerator struct {
	titles     []string
	titleIndex int
	mtx        *sync.Mutex
	shuffle    *sync.Once
}

func (r *randomItemGenerator) reset() {
	r.mtx = &sync.Mutex{}
	r.shuffle = &sync.Once{}

	r.titles = GetServices()
}

func (r *randomItemGenerator) next() item {
	if r.mtx == nil {
		r.reset()
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := item{
		title: r.titles[r.titleIndex],
	}

	r.titleIndex++
	if r.titleIndex >= len(r.titles) {
		r.titleIndex = 0
	}

	return i
}
