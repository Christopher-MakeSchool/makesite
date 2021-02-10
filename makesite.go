package main

import (
	"flag"
	"fmt"
	"github.com/foize/go.sgr"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func processFile(filePath string) {
	fileName := strings.Split(filePath, ".txt")[0]
	namedHTML := fileName + ".html"
	fileContents := readFile(filePath)

	info := dataProcessing{
		Path:    filePath,
		Name:    fileName,
		HTML:    namedHTML,
		Content: fileContents,
	}

	applyTemplate("template.tmpl", info)
}

// Create an Html File Based off the 
func applyTemplate(path string, data dataProcessing) {
	t := template.Must(template.New(path).ParseFiles(path))
	newFile, _ := os.Create(data.HTML)
	// err := t.Execute(os.Stdout, data)
	err := t.Execute(newFile, data)
	if err != nil {
		panic(err)
	}
}

func searchDirectory(dirPath string) {
	libRegEx, e := regexp.Compile("^.+\\.(txt)$")
	if e != nil {
		log.Fatal(e)
	}
	count, size := 0, 0.0

	e = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			count = count + 1
			fmt.Println(info.Name(), float64(info.Size()), float64(os.Getpagesize()))
			// size = size + float64(info.Size())
			size = size + float64(os.Getpagesize())/1000.0
			processFile(info.Name())
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
	sgr.Printf("[fg-green bold] Success! [reset] Generated [bold] %d [reset] pages (%6.1fkB total). \n", count, size)
}

func main() {
	var dirPath, filePath string
	flag.StringVar(&dirPath, "dir", "", "Directory Path")
	flag.StringVar(&filePath, "file", "", "Name or Path to a text file")
	flag.Parse()

	switch {
	case dirPath != "":
		searchDirectory(dirPath)
	case filePath != "":
		processFile(filePath)
	default:
		fmt.Print("No Option Selected")
	}
}
