package job

import (
	"net/http"

	"github.com/ikeikeikeike/prototype/injector/util"
	"github.com/labstack/echo"
)

type (
	// JobViewer is
	JobViewer struct {
		Env       *util.Env `inject:""`
		Scheduler Scheduler `inject:""`
	}
)

// JobViewer is
func (hdr *JobViewer) Show(ctx echo.Context) (err error) {
	data, err := hdr.Scheduler.Receive()
	// log.Println("[DEBUG] ", data)
	return ctx.JSON(http.StatusOK, data)
}
