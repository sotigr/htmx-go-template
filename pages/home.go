package pages

import "net/http"

type IndexPage struct{}

func (p IndexPage) Name() string {
	return "Home"
}

func (p IndexPage) URL() string {
	return "/"
}

func (p IndexPage) Template() string {
	return `
	{{define "Head"}}
		<title>Welcome htmx</title>
	{{end}}
	{{template "DocumentStart" .}}   

		<a href="/about">
			go to about
		</a>
		<h2>This is loaded from the frontend:</h2>
		<div class="bg-slate-100" hx-get="/dyn/test/test" hx-swap="innerHtml"  hx-trigger="load"></div>

		<h2>This is prerendered:</h2>
		<div class="bg-slate-100"> {{template "DynTestTest" .}}</div>
		
		<input type="text" name="text"
			hx-post="/dyn/test/form"
			hx-trigger="keyup changed delay:50ms"
			hx-target="#search-results"
			placeholder="Search..."
			class="w-full border border-black"
		>
		<div id="search-results"></div>

	{{template "DocumentEnd" .}}   
	`
}

func (p IndexPage) Props(r *http.Request) any {
	return nil
}
