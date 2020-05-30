package controller

import (
	"html/template"
	"net/http"

	"webapp/src/viewmodel"
)

type home struct  {
	homeTemplate 			*template.Template // reference to the home page template 
	
}
// register route of home // 
func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

// process request // 
func (h home) handleHome(w http.ResponseWriter, r *http.Request) {	
	vm := viewmodel.NewHome() // 
	h.homeTemplate.Execute(w, vm) // excute home template pass in the view model 

}
