package job

import (
	"time"

	"github.com/ikeikeikeike/prototype/injector/util"
	cor "github.com/ikeikeikeike/prototype/injector/util/container"
)

type (
	// Scheduler is
	Scheduler interface {
		Receive() ([]string, error)
		Exit()
		Lost()
	}

	scheduler struct {
		Env *util.Env `inject:""`

		data    []string
		receive chan *cor.Container
		exit    chan struct{}
		lost    chan struct{}
	}
)

// Receive receives data
func (s *scheduler) Receive() ([]string, error) {
	c := cor.Containers.Get()
	defer cor.Containers.Put(c)

	s.receive <- c

	b := <-c.Baggage
	defer cor.Baggages.Put(b)

	item, err := b.Item, b.Err
	return item.([]string), err
}

// Lost discards data
func (s *scheduler) Lost() {
	s.lost <- struct{}{}
}

// Exit closes goroutine loop.
func (s *scheduler) Exit() {
	s.exit <- struct{}{}
}

func (s *scheduler) lift() {
	s.data = append(s.data, "_")
}

func (s *scheduler) discharge() *cor.Baggage {
	b := cor.Baggages.Get(s.data, nil)
	s.data = make([]string, 0)
	return b
}

func (s *scheduler) run() {
	sTick := time.NewTicker(10 * time.Second)
	defer sTick.Stop()

	for {
		select {
		case <-sTick.C:
			s.lift()
		case c := <-s.receive:
			c.Baggage <- s.discharge()
		case <-s.lost:
			s.data = make([]string, 0)
		case <-s.exit:
			s.data = make([]string, 0)
			break
		}
	}
}

func bootScheduler() *scheduler {
	s := &scheduler{
		data:    make([]string, 0),
		receive: make(chan *cor.Container),
		exit:    make(chan struct{}),
		lost:    make(chan struct{}),
	}

	go s.run()

	return s
}
