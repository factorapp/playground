package pages

import (
	
    "fmt"
    "strings"
    "strconv"


	"github.com/gowasm/vecty"
)

var IndexTemplate = `<main role="main" class="container">
    <components:Nav Name="Stuff" />
    <div class="starter-template">
        <h1>Factor Web Assembly Framework</h1>
        <components:Badge BadgeType="badge-dark" Text="Danger Zone!" />
        <p class="lead">Build highly performant web applications.<br/> Using the language you love.</p>
        <p class="lead"><a href="https://github.com/factorapp/playground">Source</a></p>
        <button vecty:onclick="OnClick">Increment</button><p>{vecty-field:CountText}</p>
        <components:Alert AlertType="alert-info" Text="Hello World" />
    </div>
</main>

`
var IndexTitle =  "Page Title!"

type Index struct {
	vecty.Core
	Title string
	
    Name string
    count int
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


