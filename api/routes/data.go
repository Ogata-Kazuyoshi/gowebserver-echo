package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
)



func Data(g *echo.Group)  {
	g.GET("",func(c echo.Context) error { //Goでは / の有無で変わる。Frontからの叩きかたも要注意!!
		
		msg := map[string]string{"message": "Getできてる？"}
	// マップをJSONにエンコードしてレスポンスに書き込む
		return c.JSON(200,msg) //ステータスコードは第１引数に
	})
	g.GET("/:id",func(c echo.Context) error {
		id := c.Param("id")	
		msg := map[string]string{"message": "パスパラメータVer" ,"pathparamater:Id" : id}
	// マップをJSONにエンコードしてレスポンスに書き込む
		return c.JSON(201,msg)
	})
	g.POST("", func(c echo.Context) error {

		fmt.Println("ここまではきてるんだよなーー")

		var reqBody map[string]string

		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		return c.JSON(201,reqBody)
	})
}