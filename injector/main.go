package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/ikeikeikeike/prototype/injector/analyzer"
	"github.com/ikeikeikeike/prototype/injector/util"

	einhorn "github.com/dcu/http-einhorn"
	"github.com/facebookgo/inject"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // no need
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var env util.Env
	if err := envconfig.Process("", &env); err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to process env var: %s", err))
	}

	// routing
	e := route()
	e.Server.Addr = env.RjtrackURI

	// Injects dependecies
	var g inject.Graph

	err := g.Provide(
		&inject.Object{Value: &env},
	)
	if err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to process inject.Provide: %s", err))
	}

	analyzer.Inject(&g, e)

	if err := g.Populate(); err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to process inject.Populate: %s", err))
	}

	// Port Listen
	if einhorn.IsRunning() {
		err = einhorn.Run(e.Server, 0)
	} else {
		err = e.StartServer(e.Server)
	}

	e.Logger.Fatal(err)
}
