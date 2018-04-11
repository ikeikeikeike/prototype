package analyzer

import (
	"net/http"

	"github.com/ikeikeikeike/prototype/injector/models"
	"github.com/ikeikeikeike/prototype/injector/util"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

type (
	// Analyzer is
	Analyzer struct {
		Env    *util.Env `inject:""`
		Worker Worker    `inject:""`
	}
)

// Analyze is
func (hdr *Analyzer) Analyze(ctx echo.Context) (err error) {
	token := "unko "

	hdr.Worker.Send(token)

	users, err := models.PilotsG().All()
	pp.Println(users, err)

	return ctx.JSON(http.StatusOK, token)
}
