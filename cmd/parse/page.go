package main

var pageTemplate = `package pages

import (
	{{ .Imports }}

	"github.com/gowasm/vecty"
)

var {{.Name}}Template = {{.Template}}
var {{.Name}}Title = {{.Title}}

type {{.Name}} struct {
	vecty.Core
	Title string
	{{ .Parameters}}
}

func (c *{{.Name}}) GetTitle() string {
	if c.Title != "" {
		return c.Title
	}
	return {{.Name}}Title
}
func (c *{{.Name}}) Template() string {
	return {{ .Name}}Template
}


{{ .Funcs}}
`
