// This file was created with https://github.com/factorapp/factor
// using https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

func (p *Badge) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("badge", p.BadgeType),
		),
		vecty.Text(p.Text),
	)
}
