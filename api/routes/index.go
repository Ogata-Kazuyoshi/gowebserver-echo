package routes

import "github.com/labstack/echo/v4"

func Router(g *echo.Group)  {
	dataGroup :=g.Group("/data")
	Data(dataGroup)

	authGroup :=g.Group("/auth")
	Auth(authGroup)
}