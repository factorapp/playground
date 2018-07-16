// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

func (p *Alert) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("alert", p.AlertType, "alert-dismissible", "fade", "show"),
			vecty.Attribute("role", "alert"),
		),
		vecty.Text(p.Text),
		elem.Button(
			vecty.Markup(
				prop.Type(prop.TypeButton),
				vecty.Class("close"),
				vecty.Data("dismiss", "alert"),
				vecty.Attribute("aria-label", "Close"),
			),
			elem.Span(
				vecty.Markup(
					vecty.Attribute("aria-hidden", "true"),
				),
				vecty.Text("x"),
			),
		),
	)
}
