package sync

import "sync"

type Counter struct {
	// sync.Mutex embedded into the struct, this looks nice but while
	// programming is a hugely subjective discipline, this is bad and wrong.
	// Sometimes people forget that embedding types means the methods of that
	// type becomes part of the public interface; and you often will not want that.
	// Remember that we should be very careful with our public APIs, the moment we
	// make something public is the moment other code can couple themselves to it.
	// We always want to avoid unnecessary coupling.

	sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
