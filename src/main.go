package main

import (
	"net/http"

	"webapp/src/controller"
	"webapp/src/helpers"
)


func main() {

	// SETUP THE CONTROLLER //
	templates := helpers.PopulateTemplates()

	controller.Startup(templates)
	
	http.ListenAndServe(":8000", nil)	
}
