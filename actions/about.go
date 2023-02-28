package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"net/http"
)

func About(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON("root about info"))
}

func VirtualHost1About(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON("virtual host 1 about info"))
}

func VirtualHost2About(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON("virtual host 2 about info"))
}

func G1About(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON("g1 about info"))
}

func G2About(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.JSON("g2 about info"))
}
