package main

import (
	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic-github/internal/capabilities"
)

func main() {
	e := echo.New()

	s := &capabilties.PluginService{}
	e.POST("/get-capabilities", s.GetCapabilities)

	e.Logger.Fatal(e.Start(":8080"))
}
