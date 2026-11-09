package main

import (
	"fmt"
	"sync"
)

type Span struct {
	size      int
	allocated bool
}

type mheap struct {
	spans []*Span
	lock  sync.Mutex
}

type mcentral struct {
	sizeSpans []*Span
	lock sync.Mutex
}

type mcache struct {
	localSpan []*Span
}

func NewHeap(size int) *mheap {
	h := &mheap{}

	for i := 0; i < size; i++ {
		h.spans = append(h.spans, &Span{size: 1 + i, allocated: false})
	}	
	
	return h
}

func (h *mheap) getSpan(size int) *Span {
	h.lock.Lock()
	defer h.lock.Unlock()

	for _, span := range h.spans {
		if !span.allocated && span.size >= size {
			span.allocated = true
			return span
		}
	}

	return nil
}

func (c *mcentral) getSpan(size int) *Span {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, span := range c.sizeSpans {
		if !span.allocated && span.size >= size {
			span.allocated = true
			return span
		}
	}

	return nil
}

func (c *mcache) getSpan(size int) *Span {
	for _, span := range c.localSpan {
		if !span.allocated && span.size >= size {
			span.allocated = true
			return span
		}
	}

	return nil
}

func main() {
	heap := NewHeap(10)
	mcentral := &mcentral{
		sizeSpans: heap.spans,
	}
	mcache := &mcache{}

	reqSpan := mcache.getSpan(5)

	if reqSpan == nil {
		reqSpan = mcentral.getSpan(5)
	}
	if reqSpan == nil {
		reqSpan = heap.getSpan(5)
	}
	if reqSpan == nil {
		panic("out of memory")
	}

	fmt.Println("Allocated span of size:", reqSpan.size)
}