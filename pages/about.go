package pages

import "net/http"

type AboutPage struct{}

func (p AboutPage) Name() string {
	return "About"
}

func (p AboutPage) URL() string {
	return "/about"
}

func (p AboutPage) Template() string {
	return ` 
	{{template "DocumentStart" .}}   
	about
	{{template "DocumentEnd" .}}    
	`
}

func (p AboutPage) Props(r *http.Request) any {
	return nil
}
