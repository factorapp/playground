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

func (p *Components) Render() vecty.ComponentOrHTML {
	vecty.SetTitle(p.GetTitle())
	return elem.Body(
		elem.Main(
			vecty.Markup(
				vecty.Attribute("role", "main"),
				vecty.Class("container"),
			),
			&components.Nav{Name: "Factor"},
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
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Attribute("name", "examples"),
				),
				elem.Heading2(
					vecty.Markup(
						prop.ID("examples"),
					),
					vecty.Text("Examples"),
				),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Attribute("name", "alerts"),
				),
				elem.Paragraph(
					vecty.Text("Alerts"),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("bd-example"),
				),
				&components.Alert{
					AlertType: "alert-primary",
					Text:      "A simple primary alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-secondary",
					Text:      "A simple secondary alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-success",
					Text:      "A simple success alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-danger",
					Text:      "A simple danger alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-warning",
					Text:      "A simple warning alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-info",
					Text:      "A simple info alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-light",
					Text:      "A simple light alert—check it out!",
				},
				&components.Alert{
					AlertType: "alert-dark",
					Text:      "A simple dark alert—check it out!",
				},
			),
		),
	)
}
