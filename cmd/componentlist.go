package main

import (
	"server/core"
	dyntest "server/dynamic/test"
	"server/pages"
)

func GetComponents() *[]core.Component {
	return &[]core.Component{
		pages.Document{},
		pages.IndexPage{},
		pages.AboutPage{},
		dyntest.TestTest{},
		dyntest.TestTestForm{},
	}
}
