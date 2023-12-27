package dyntest

import (
	"net/http"
)

type TestTestForm struct{}

func (p TestTestForm) Name() string {
	return "DynTestTestForm"
}

func (p TestTestForm) URL() string {
	return "/dyn/test/form"
}

func (p TestTestForm) Template() string {
	return `
	the text is: <span class="text-red-900">{{.Props.Text}}</span>
	`
}

func (p TestTestForm) Props(r *http.Request) any {
	r.ParseForm()           // Parses the request body
	x := r.Form.Get("text") // x will be "" if parameter is not set
	return struct {
		Text string
	}{Text: x}
}
