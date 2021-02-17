package cmd

import (
	"html/template"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/chrisbarnes2000/makesite/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fileCommand)
}

var fileCommand = &cobra.Command{
	Use:   "file [filePath]",
	Short: "txt to html",
	Long:  "Convert The Provided File From .txt or .md to .html",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		parsedPath := processFilePath(filePath)
		fileLocation := parsedPath[1]
		fileName := parsedPath[2]
		fileExtension := parsedPath[3]
		namedHTML := fileName + ".html"

		fileInfo := &models.DataProcessing{
			Path: fileLocation,
			Name: fileName,
			HTML: namedHTML,
		}

		if fileExtension == ".txt" {
			fileInfo.Content = processTxtFile(filePath)
		} else if fileExtension == ".md" {
			fileInfo.Content = processMdFile(filePath)
		}

		applyTemplate("template.tmpl", fileInfo)
	},
}

func processFilePath(filePath string) []string {
	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	parsedPath := re.FindStringSubmatch(filePath)
	return parsedPath
}

func processTxtFile(filePath string) string {
	return readFile(filePath)
}

func processMdFile(filePath string) string {
	// Original Plan: Parse Line By Line and Convert Markdown Syntax to HTML
	// Better Plan: Use 3rd party https://github.com/gomarkdown/markdown
	fileContents := readFile(filePath)
	// switch fileContents {
	// case "#":
	//     return "<h1></h1>"
	// case "##":
	//     return "<h2></h2>"
	// case "###":
	//     return "<h3></h3>"
	// }
	return fileContents + "\nComplete Markdown Conversion"
}

// Read a file given its path/name
func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

// Create an Html File Based off the provided template and processed data
func applyTemplate(path string, data *models.DataProcessing) {
	htmlFolder := "./html/"
	tmplFolder := "./tmpl/"
	// Create a New HTML File to Apply the Template to
	newFile, _ := os.Create(htmlFolder + data.HTML)
	
	// Create a temparay html template file by parsing the tmpl Tempalte at `path`
	t := template.Must(template.New(path).ParseGlob(tmplFolder+path))
	
	// err := t.Execute(os.Stdout, data)
	err := t.Execute(newFile, data)
	if err != nil {
		panic(err)
	}
}
