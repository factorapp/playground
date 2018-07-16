package components

import (
	


	"github.com/gowasm/vecty"
)

var BadgeTemplate = `<span class="badge {vecty-field:BadgeType}">{vecty-field:Text}</span>`

type Badge struct {
	vecty.Core
	
    Text string
    BadgeType string


}

func (c *Badge) Template() string {
	return BadgeTemplate
}





