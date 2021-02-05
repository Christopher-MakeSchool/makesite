package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"regexp"
	"path/filepath"
)

// Stores the contents of file as a string
type dataProcessing struct {
	Path    string
	Name    string
	HTML    string
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
func renderTemplate(path string, data dataProcessing) {
	t := template.Must(template.New(path).ParseFiles(path))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	newFile, _ := os.Create(data.HTML)
	t.Execute(newFile, data)
	fmt.Print("Saved File: ", data.HTML)
}

func main() {
	var dirPath, filePath string
	flag.StringVar(&dirPath, "dir", "", "Directory Path")
	flag.StringVar(&filePath, "file", "", "Name or Path to a text file")
	flag.Parse()

	// fmt.Print(dirPath, filePath)

	switch {
	case dirPath != "":
		libRegEx, e := regexp.Compile("^.+\\.(txt)$")
		if e != nil {
				log.Fatal(e)
		}

		e = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
				if err == nil && libRegEx.MatchString(info.Name()) {
						println(info.Name())
				}
				return nil
		})
		if e != nil {
				log.Fatal(e)
		}

		// files, err := ioutil.ReadDir(dirPath)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// for _, file := range files {
		// 	fmt.Println(file.Name())
		// }

	case filePath != "":
		fileName := strings.Split(filePath, ".txt")[0]
		namedHTML := fileName + ".html"
		fileContents := readFile(filePath)

		info := dataProcessing{
			Path:    filePath,
			Name:    fileName,
			HTML:    namedHTML,
			Content: fileContents,
		}
		// fmt.Println(info)

		renderTemplate("template.tmpl", info)

	default:
		fmt.Print("No Option Selected")
	}
}
