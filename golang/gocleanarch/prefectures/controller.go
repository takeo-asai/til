package prefectures

import (
	"github.com/labstack/echo"
	"net/http"

	"sampleapp/prefectures/find"
)

// Controller :
type Controller struct {
	Interactor *find.Interactor
}

// NewController :
func NewController(it *find.Interactor) *Controller {
	return &Controller{it}
}

// Index :
func (ctrl *Controller) Index(c echo.Context) error {
	req := &find.Request{} // no params

	resp, err := ctrl.Interactor.Handle(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp.Prefectures)
}
