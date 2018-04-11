package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ikeikeikeike/prototype/injector/analyzer"
	"github.com/ikeikeikeike/prototype/injector/job"
	"github.com/ikeikeikeike/prototype/injector/util"
	"github.com/volatiletech/sqlboiler/boil"

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

	// db
	db, err := sql.Open("mysql", env.DSN)
	if err != nil {
		panic(fmt.Sprintf("It was unable to connect to the DB.\n%s\n", err))
	}
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	defer db.Close()

	// models
	boil.DebugMode = !env.IsProd()
	boil.SetDB(db)

	// routing
	e := route()
	e.Server.Addr = env.URI

	// Injects dependecies
	var g inject.Graph

	err = g.Provide(
		&inject.Object{Value: &env},
	)
	if err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to process inject.Provide: %s", err))
	}

	job.Inject(&g, e)
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
