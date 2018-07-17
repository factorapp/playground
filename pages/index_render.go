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
					vecty.Text("Factor Web Assembly Framework"),
				),
				&components.Badge{
					BadgeType: "badge-dark",
					Text:      "Danger Zone!",
				},
				elem.Paragraph(
					vecty.Markup(
						vecty.Class("lead"),
					),
					vecty.Text("Build highly performant web applications."),
					elem.Break(),
					vecty.Text("Using the language you love."),
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
				&components.Alert{
					AlertType: "alert-info",
					Text:      "Hello World",
				},
			),
		),
	)
}
