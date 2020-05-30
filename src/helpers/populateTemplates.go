package helpers

import (
	"os"
	"html/template"
	"io/ioutil"
)




// PopulateTemplates is responsible for traversing the base templates
// pull layout then clone each individual content section
// Loads whats in the subtempales  and content
func PopulateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))             // loading the template
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html")) // loading subtemplates
	dir, err := os.Open(basePath + "/content")   
	// check error
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1) // read conetnt in the and get file identifies
	// check error
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	// loop over the files found in the content folder
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name()) // open up each file
		// check for error
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}

		content, err := ioutil.ReadAll(f) // read contents in file
		// check error
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close() // close file
		// Create template instance //

		tmpl := template.Must(layout.Clone()) // clone layout template and child and clone into a variable called templ
		_, err = tmpl.Parse(string(content))  // parse the templ
		// check error
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl // add tempale into a result map.
	}
	return result // return result map
}
