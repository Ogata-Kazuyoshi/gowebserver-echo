package api

import (
	"net/http"
	"server/api/routes"

	"github.com/labstack/echo/v4"
)

func StartMainServer()  {
	server := echo.New()

	apiGroup := server.Group("/api/v1")
	routes.Router(apiGroup)

	// ルートディレクトリにアクセスが来た際に、distフォルダのindex.htmlを返却
	server.GET("/", func(c echo.Context) error {
		return c.File("dist/index.html")
	})

	// /testのエンドポイントにアクセスが来た際に、ルートディレクトリにリダイレクト
	server.GET("/test", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/")
	})


	server.Start(":4242")
}