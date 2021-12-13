package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func findHtmlFilesInDir(blogpath string) []string {
	var ret []string
	files, err := ioutil.ReadDir(blogpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			ret = append(ret, file.Name())
		}
	}
	return ret
}

func copyBlogHtmlFiles(blogpath string, outpath string) {
	// Ignoring errors for `mkdir -p` equivalent.  How do do better?
	os.Mkdir(outpath, 0755)
	for _, file := range findHtmlFilesInDir(blogpath) {
		content, err := ioutil.ReadFile(blogpath + "/" + file)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(outpath+"/"+file, []byte(content), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	var blogpath, outpath string
	flag.StringVar(&blogpath, "b", "", "Specify path to Org export / blog files")
	flag.StringVar(&outpath, "o", "/tmp/blog", "Specify path for target directory")
	flag.Parse()
	if blogpath == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	copyBlogHtmlFiles(blogpath, outpath)
	fmt.Println("OK")
}
