package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func listHtmlFilesInDir(blogpath string) {
	files, err := ioutil.ReadDir(blogpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			fmt.Println(strings.Replace(file.Name(),
				".html", "", 1))
		}
	}
}

func main() {
	var blogpath string
	flag.StringVar(&blogpath, "b", "", "Specify path to Org export / blog files.")
	flag.Parse()
	if blogpath == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	listHtmlFilesInDir(blogpath)
}
