package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Request     *http.Request
	Props       any
	PayloadJson string
	Page        string
}

func main() {

	Components := GetComponents()

	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	templates := template.Must(template.New(""), nil)
	for _, component := range *Components {
		c := component
		name := c.Name()
		tmp := templates.New(name)
		tmp, err := tmp.Parse(c.Template())
		url := c.URL()
		if url != "" {
			mux.HandleFunc(c.URL(), func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Header().Set("CharSet", "utf-8")

				if err != nil {
					io.WriteString(w, err.Error())
					return
				}

				props := c.Props(r)

				if err != nil {
					io.WriteString(w, err.Error())
					return
				}

				tmp.Execute(w, TemplateData{
					Props:   props,
					Page:    name,
					Request: r,
				})
			})
		}
	}

	fmt.Println("Listening")

	http.ListenAndServe("0.0.0.0:3000", mux)

}
