package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	// aboutController about
)


// register static resource handlers and pull the template from the template map 
// pass into the home and shop controller register router method 
func Startup(templates map[string] *template.Template) {

	homeController.homeTemplate = templates["home.html"] // assign template field base on what in the map
	// homeController.loginTemplate = templates["about.html"]

	homeController.registerRoutes() // register route 
	// aboutController.registerRoutes()



	// Serve static folder //	
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}