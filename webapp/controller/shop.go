package controller

import (
	"html/template"
	"net/http"

	"lssGo/webapp/viewmodel"
)

type shop struct {
	shopTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	//http.HandleFunc("/", s.handleHome)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewShop()
	s.shopTemplate.Execute(w, vm)
}
