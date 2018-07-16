package components

import (
	


	"github.com/gowasm/vecty"
)

var AlertTemplate = `<div class="alert {vecty-field:AlertType} alert-dismissible fade show" role="alert">{vecty-field:Text}<button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">x</span></button></div>`

type Alert struct {
	vecty.Core
	
    Text string
    AlertType string


}

func (c *Alert) Template() string {
	return AlertTemplate
}





