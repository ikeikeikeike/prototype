package job

import (
	"fmt"

	"github.com/facebookgo/inject"
	"github.com/labstack/echo"
)

// Inject injects dependencies
func Inject(g *inject.Graph, e *echo.Echo) {
	var jv JobViewer

	// inject
	err := g.Provide(
		&inject.Object{Value: &jv},
		&inject.Object{Value: bootScheduler()},
	)
	if err != nil {
		panic(fmt.Sprintf("[ERROR] Failed to process injection: %s", err))
	}

	// routes
	gp := e.Group("/job")
	gp.GET("", jv.Show)
}
