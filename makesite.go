package main

import (
	"fmt"
	"os"
	// "flag"
	"html/template"
	"io/ioutil"
)

// Stores the contents of file as a string
type dataFile struct {
	Content string
}

// Read a file given its path/name
func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

// Create a template based of a given file
func renderTemplate(path, outputFile string, data dataFile) {
	t := template.Must(template.New(path).ParseFiles(path))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	newFile, _ := os.Create(outputFile)
	t.Execute(newFile, data)
	fmt.Print("Saved File: ", outputFile)
}

// // Fill a template with the given content
// func writeHTML(content, newFileName string) {
// 	bytesToWrite := []byte(content)
// 	err := ioutil.WriteFile(newFileName, bytesToWrite, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
	// fmt.Print("Saved File", newFileName)
// }

func main() {
	// fmt.Println("Running Main Function")
	// fmt.Println("\nv1.0 Reg#2: Read the contents of 'first-post.txt' ")
	fileContents := dataFile{readFile("first-post.txt")}

	// fmt.Println("v1.0 Reg#3&4: Edit and Print the provided html template 'template.tmpl' with the contents of 'first-post.txt' ")
	// fmt.Println("\nv1.0 Reg#5: Write the html template to a file named 'first-post.html' ")
	renderTemplate("template.tmpl", "first-post.html", fileContents)
	
}
