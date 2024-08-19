package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a simple template string with placeholders
	tmpl := `Hello, {{.Name}}! Welcome to {{.Location}}.`

	// Parse the template string
	t, err := template.New("greeting").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Create a data structure to inject into the template
	data := struct {
		Name     string
		Location string
	}{
		Name:     "Alice",
		Location: "Wonderland",
	}

	// Execute the template, writing the output to os.Stdout
	err = t.Execute(os.Stdout, data)
	// Hello, Alice! Welcome to Wonderland.%   
	if err != nil {
		panic(err)
	}
}
