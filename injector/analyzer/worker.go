package analyzer

import (
	"log"

	"github.com/ikeikeikeike/prototype/injector/util"
)

type (
	// Worker manifests encoder public interface
	Worker interface {
		Send(data string)
	}

	// throwWorker works as background worker which encodes media stream to dash
	throwWorker struct {
		Env  *util.Env `inject:""`
		data chan string
	}
)

func (e *throwWorker) Send(data string) {
	e.data <- data
}

func (e *throwWorker) run() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ERROR] A throwWorker panic occured, this routine will restarted", r)
			go e.run()
		}
	}()

	for {
		e.encode(<-e.data)
	}
}

func (e *throwWorker) encode(data string) {
	log.Println("[INFO] successfuly encode: ", data)
}

func bootThrowWorker() *throwWorker {
	s := &throwWorker{
		data: make(chan string, 100000),
	}

	go s.run()

	return s
}
