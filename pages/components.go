package pages

import (
	
    "fmt"
    "strings"
    "strconv"


	"github.com/gowasm/vecty"
)

var ComponentsTemplate = `<main role="main" class="container">
    <components:Nav Name="Factor" />
    <div class="starter-template">
        <h1>Factor Web Assembly Framework</h1>
        <components:Badge BadgeType="badge-dark" Text="Danger Zone!" />
        <p class="lead">Build highly performant web applications.<br/> Using the language you love.</p>
        <p class="lead"><a href="https://github.com/factorapp/playground">Source</a></p>
        <button vecty:onclick="OnClick">Increment</button><p>{vecty-field:CountText}</p>
    </div>
<a name="examples"><h2 id="examples">Examples</h2></a>

<a name="alerts"><p>Alerts</p></a>

<div class="bd-example">

<components:Alert AlertType="alert-primary" Text="A simple primary alert—check it out!" />
<components:Alert AlertType="alert-secondary" Text="A simple secondary alert—check it out!" />
<components:Alert AlertType="alert-success" Text="A simple success alert—check it out!" />

<components:Alert AlertType="alert-danger" Text="A simple danger alert—check it out!" />
<components:Alert AlertType="alert-warning" Text="A simple warning alert—check it out!" />
<components:Alert AlertType="alert-info" Text="A simple info alert—check it out!" />
<components:Alert AlertType="alert-light" Text="A simple light alert—check it out!" />
<components:Alert AlertType="alert-dark" Text="A simple dark alert—check it out!" />
</div>
</main>

`
var ComponentsTitle =  "Components"

type Components struct {
	vecty.Core
	Title string
	
    Name string
    count int
    CountText string

}

func (c *Components) GetTitle() string {
	if c.Title != "" {
		return c.Title
	}
	return ComponentsTitle
}
func (c *Components) Template() string {
	return ComponentsTemplate
}


 
    func (p *Components) LowerName() string {
        return strings.ToLower(p.Name)
    }

func (p *Components) OnClick(e *vecty.Event) {
	fmt.Println("Someone clicked on me", e.Target)
	p.count++
	p.CountText = "Click Count: " + strconv.Itoa(p.count)

	vecty.Rerender(p)
}


