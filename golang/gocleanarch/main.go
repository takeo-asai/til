package main

import (
	"github.com/labstack/echo"

	"sampleapp/prefectures"
	"sampleapp/prefectures/find"
	"sampleapp/datastore"
)

func main() {
	e := echo.New()

	ctrl := prefectures.NewController(
		find.NewInteractor(
			datastore.NewPrefectureRepositoryImpl()))
	e.GET("/prefectures", ctrl.Index)
	e.Logger.Fatal(e.Start(":1323"))
}
