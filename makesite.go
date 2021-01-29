package main

import (
	"os"
	"fmt"
	// "flag"
	"io/ioutil"
	"html/template"
)

func main() {
	fmt.Println("Running Main Function")
	fmt.Println("\nv1.0 Reg#2: Read the contents of 'first-post.txt' ")
	file_contents := data_file{readFile("first-post.txt")}
	// fmt.Println("\n", file_contents)
	fmt.Println("\nv1.0 Reg#3: Edit the provided html template 'template.tmpl' ")
	renderTemplate("template.tmpl", "template.tmpl", file_contents)
	// fmt.Println("\n", )

}

func readFile(file_name string) string {
	fileContents, err := ioutil.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

type data_file struct {
	Content string
}

func renderTemplate(output_template_name, path string, data data_file) *template.Template {
	t := template.Must(template.New(output_template_name).ParseFiles(path))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	return t
}