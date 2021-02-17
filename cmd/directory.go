package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/foize/go.sgr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dirCommand)
}

var dirCommand = &cobra.Command{
	Use:   "dir",
	Short: "List all text files",
	Long:  "List all text files in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		searchDirectory(args[0], args[1])
	},
}

func searchDirectory(dirPath, fileType string) {
	count, size, regex := 0, 0.0, ""

	if fileType == ".txt" {
		regex = "^.+\\.(txt)$"
	} else if fileType == ".md" {
		regex = "^.+\\.(md)$"
	}
	libRegEx, e := regexp.Compile(regex)
	if e != nil {
		log.Fatal(e)
	}

	// Walk through the directory looking for files the match our regex parse above
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
