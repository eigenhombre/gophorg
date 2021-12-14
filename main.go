// gophorg is an experimental reboot of my blog in (*gasp*) Go.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func findHTMLFilesInDir(blogpath string) ([]string, error) {
	var ret []string
	files, err := ioutil.ReadDir(blogpath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			ret = append(ret, file.Name())
		}
	}
	return ret, nil
}

func processHTMLFile(blogpath, outpath, file string) error {
	content, err := ioutil.ReadFile(blogpath + "/" + file)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outpath+"/"+file, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func copyBlogHTMLFiles(blogpath string, outpath string) (int, error) {
	var err error
	if err = os.MkdirAll(outpath, 0755); err != nil {
		return 0, err
	}
	files, err := findHTMLFilesInDir(blogpath)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		err := processHTMLFile(blogpath, outpath, file)
		if err != nil {
			return 0, err
		}
	}
	return len(files), nil
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
	numfiles, err := copyBlogHTMLFiles(blogpath, outpath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Processed %d exported files.\n", numfiles)
}
