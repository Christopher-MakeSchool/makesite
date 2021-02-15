package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chrisbarnes2000/makesite/models"
	"github.com/foize/go.sgr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dirCommand)
	rootCmd.AddCommand(fileCommand)
}

var dirCommand = &cobra.Command{
	Use:   "dir",
	Short: "List all text files",
	Long:  "List all text files in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		searchDirectory(".")
	},
}

var fileCommand = &cobra.Command{
	Use:   "file [filePath] [templatePath]",
	Short: "txt to html",
	Long:  "Convert The Provided File From Txt to Html",
	Run: func(cmd *cobra.Command, args []string) {
		processFile(args[0], args[1])
	},
}

// Read a file given its path/name
func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func processFile(filePath, templatePath string) {
	fileName := strings.Split(filePath, ".txt")[0]
	namedHTML := fileName + ".html"
	fileContents := readFile(filePath)

	info := &models.DataProcessing{
		Path:    filePath,
		Name:    fileName,
		HTML:    namedHTML,
		Content: fileContents,
		ID:      "Onf6Bvx0",
	}

	applyTemplate("flyer-template.tmpl", info)
	// applyTemplate("template.tmpl", info)
}

// Create an Html File Based off the provided template and processed data
func applyTemplate(path string, data *models.DataProcessing) {
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
			fmt.Println(info.Name())
			size = size + float64(os.Getpagesize())/1000.0
			// if *&Verbose {
			// 	processFile(info.Name())
			// }
		}
		return nil
	})
	if e != nil {
		log.Fatal("Error walking through the Directory", e)
	}
	if *&Verbose {
		sgr.Printf("[fg-green bold] Success! [reset] Generated [bold] %d [reset] pages (%6.1fkB total). \n", count, size)
	}
}
