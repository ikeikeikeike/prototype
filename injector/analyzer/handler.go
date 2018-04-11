package analyzer

import (
	"net/http"

	"github.com/ikeikeikeike/prototype/injector/util"
	"github.com/labstack/echo"
)

type (
	// Analyzer is
	Analyzer struct {
		Env     *util.Env `inject:""`
		Encoder Encoder   `inject:""`
	}
)

// Analyze is
func (hdr *Analyzer) Analyze(ctx echo.Context) (err error) {
	token := "unko unko unko"

	hdr.Encoder.Send(token)

	return ctx.JSON(http.StatusOK, token)
}
