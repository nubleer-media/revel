package controllers

import "github.com/nubleer/revel"

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}
