package container

import "sync"

// Container TODO later
type Container struct {
	Baggage chan *Baggage
}

var defcontainer = &Container{
	Baggage: make(chan *Baggage),
}

func (c *Container) reset() {
	c.Baggage = defcontainer.Baggage
}

type containerPool struct {
	sync.Pool
}

// Get container from sync.Pool
func (cp *containerPool) Get() *Container {
	return cp.Pool.Get().(*Container)
}

// Put container to sync.Pool
func (cp *containerPool) Put(c *Container) {
	c.reset()
	cp.Pool.Put(c)
}

// Containers is TODO later
var Containers = &containerPool{
	Pool: sync.Pool{New: func() interface{} {
		return &Container{Baggage: make(chan *Baggage)}
	}},
}

// Baggage is TODO later
type Baggage struct {
	Item interface{}
	Err  error
}

var defbaggage = &Baggage{}

func (b *Baggage) reset() {
	b.Item = defbaggage.Item
	b.Err = defbaggage.Err
}

type baggagePool struct {
	sync.Pool
}

// Get baggage from sync.Pool
func (bp *baggagePool) Get(item interface{}, err error) *Baggage {
	b := bp.Pool.Get().(*Baggage)
	b.Item = item
	b.Err = err

	return b
}

// Put baggage to sync.Pool after default to value
func (bp *baggagePool) Put(b *Baggage) {
	b.reset()
	bp.Pool.Put(b)
}

// Baggages is TODO later
var Baggages = &baggagePool{
	Pool: sync.Pool{New: func() interface{} {
		return &Baggage{}
	}},
}
