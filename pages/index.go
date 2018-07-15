package pages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gowasm/vecty"
)

var IndexTemplate = `<main role="main" class="container">
    <components:Nav Name="Stuff" />
    <div class="starter-template">
        <h1>Bootstrap starter template</h1>
        <p class="lead">Use this document as a way to quickly start any new project.<br/> All you get is this text and a mostly barebones HTML document.</p>
        <p class="lead"><a href="https://github.com/factorapp/playground">Source</a></p>
        <button vecty:onclick="OnClick">Increment</button><p>{vecty-field:CountText}</p>
    </div>
</main>

`
var IndexTitle = "Page Title!"

type Index struct {
	vecty.Core
	Title string

	Name      string
	count     int
	CountText string
}

func (c *Index) GetTitle() string {
	if c.Title != "" {
		return c.Title
	}
	return IndexTitle
}
func (c *Index) Template() string {
	return IndexTemplate
}

func (p *Index) LowerName() string {
	return strings.ToLower(p.Name)
}

func (p *Index) OnClick(e *vecty.Event) {
	fmt.Println("Someone clicked on me", e.Target)
	p.count++
	p.CountText = "Click Count: " + strconv.Itoa(p.count)

	vecty.Rerender(p)
}
