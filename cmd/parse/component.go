package main

var compTemplate = `package components

import (
	{{ .Imports }}

	"github.com/gowasm/vecty"
)

var {{.Name}}Template = {{.Template}}

type {{.Name}} struct {
	vecty.Core
	{{ .Parameters}}
}

func (c *{{.Name}}) Template() string {
	return {{ .Name}}Template
}


{{ .Funcs}}
`
