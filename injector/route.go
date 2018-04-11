package main

import (
	"strings"

	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func route() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Pre(md.RemoveTrailingSlash())

	e.Use(md.LoggerWithConfig(md.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			h := ctx.Request().Host
			if strings.HasPrefix(h, "localhost") {
				return false
			} else if strings.HasPrefix(h, "127.0.0.1") {
				return false
			}

			e.Logger.SetLevel(log.WARN)
			return true
		},
	}))

	e.Use(md.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	return e
}
