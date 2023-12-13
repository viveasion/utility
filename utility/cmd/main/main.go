package main

import (
	"github.com/labstack/echo/v4"
	"utility/cmd"
	. "utility/common/config"
	"utility/common/logger"
)

func main() {
	logger.Init(ROOT+"/log", ConfigFile.MustValue("global", "log_level", "DEBUG"))
	e := echo.New()
	cmd.serveStatic(e)
}
