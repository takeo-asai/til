package main

import (
	"github.com/labstack/echo"

	"sampleapp/prefectures"
	"sampleapp/prefectures/find"
)

func main() {
	e := echo.New()

	ctrl := prefectures.NewController(
		find.NewInteractor(
			find.NewRepositoryImpl()))
	e.GET("/prefectures", ctrl.Index)
	e.Logger.Fatal(e.Start(":1323"))
}
