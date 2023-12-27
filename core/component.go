package core

import "net/http"

type Component interface {
	Name() string
	URL() string
	Template() string
	Props(r *http.Request) any
}
