package analyzer

import (
	"log"

	"github.com/ikeikeikeike/prototype/injector/util"
	"github.com/k0kubun/pp"
)

type (
	// Encoder manifests encoder public interface
	Encoder interface {
		Send(data string)
	}

	// flacEncoder works as background worker which encodes media stream to flac
	flacEncoder struct {
		Env  *util.Env `inject:""`
		data chan string
	}
)

func (e *flacEncoder) Send(data string) {
	e.data <- data
}

func (e *flacEncoder) run() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ERROR] A flacEncoder panic occured, this routine will restarted", r)
			go e.run()
		}
	}()

	for {
		e.encode(<-e.data)
	}
}

func (e *flacEncoder) encode(data string) {
	log.Println("[INFO] successfuly encode: ", data)

	pp.Println(e)
}

func bootFlacEncoder() *flacEncoder {
	s := &flacEncoder{
		data: make(chan string, 100000),
	}

	go s.run()

	return s
}
