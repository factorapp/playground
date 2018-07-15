// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package pages

import (
	components "github.com/factorapp/playground/components"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/event"
	"github.com/gowasm/vecty/prop"
)

func (p *Index) Render() vecty.ComponentOrHTML {
	vecty.SetTitle(p.GetTitle())
	return elem.Body(
		elem.Main(
			vecty.Markup(
				vecty.Attribute("role", "main"),
				vecty.Class("container"),
			),
			&components.Nav{Name: "Stuff"},
			elem.Div(
				vecty.Markup(
					vecty.Class("starter-template"),
				),
				elem.Heading1(
					vecty.Text("Bootstrap starter template"),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("lead"),
					),
					vecty.Text("Use this document as a way to quickly start any new project."),
					elem.Break(),
					vecty.Text("All you get is this text and a mostly barebones HTML document."),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("lead"),
					),
					elem.Anchor(
						vecty.Markup(
							prop.Href("https://github.com/factorapp/playground"),
						),
						vecty.Text("Source"),
					),
				),
				elem.Button(
					vecty.Markup(
						event.Click(p.OnClick),
					),
					vecty.Text("Increment"),
				),
				elem.Paragraph(
					vecty.Text(p.CountText),
				),
			),
		),
	)
}
