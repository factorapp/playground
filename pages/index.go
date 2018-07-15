package pages

import (
	
    "strings"


	"github.com/gowasm/vecty"
)

var IndexTemplate = `<main role="main" class="container">
    <components:Nav Name="Stuff" />
    <div class="starter-template">
        <h1>Bootstrap starter template</h1>
        <p class="lead">Use this document as a way to quickly start any new project.<br/> All you get is this text and a mostly barebones HTML document.</p>
    </div>
</main>

`
var IndexTitle =  "Page Title!"

type Index struct {
	vecty.Core
	Title string
	
    Name string

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

