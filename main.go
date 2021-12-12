package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// FIXME: Decomplect this
func readHtmlFilesInDir(blogpath string) {
	files, err := ioutil.ReadDir(blogpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			fmt.Print(strings.Replace(file.Name(),
				".html", "", 1) + " ")
			content, err := ioutil.ReadFile(blogpath + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(len(content))
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
	readHtmlFilesInDir(blogpath)
}
