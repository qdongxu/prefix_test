package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"net/http"
)

func About(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON("about info"))
}
