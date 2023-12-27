package dyntest

import (
	"net/http"
)

type TestTest struct{}

func (p TestTest) Name() string {
	return "DynTestTest"
}

func (p TestTest) URL() string {
	return "/dyn/test/test"
}

func (p TestTest) Template() string {
	return `
	<div> yo whats up htmx </div>
	`
}

func (p TestTest) Props(r *http.Request) any {
	return nil
}
