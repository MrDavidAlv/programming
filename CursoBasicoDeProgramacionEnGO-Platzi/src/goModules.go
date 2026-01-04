package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func moduloEcho() {
	// instanciar echo module
	e := echo.New()
	// definir ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// iniciar servidor
	e.Logger.Fatal(e.Start(":1323"))
}
