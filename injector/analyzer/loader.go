package analyzer

import (
	"fmt"

	"github.com/facebookgo/inject"
	"github.com/labstack/echo"
)

// Inject injects dependencies and Route set handler into mux
func Inject(g *inject.Graph, e *echo.Echo) {
	var ana Analyzer

	// inject
	err := g.Provide(
		&inject.Object{Value: bootThrowWorker()},
		&inject.Object{Value: &ana},
	)
	if err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to process injection: %s", err))
	}

	// routes
	gp := e.Group("/analyze")
	// g.Use(md.BasicAuth(basicAuth))
	gp.GET("", ana.Analyze)
}
