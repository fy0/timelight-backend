package main

import (
	"timelight-backend/api"
	"timelight-backend/core"
	"timelight-backend/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	core.LoggerInit()
	model.DBInit()
	defer model.GetDB().Close()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api.Bind(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
