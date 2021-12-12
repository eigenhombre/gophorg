package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var blogpath string
	flag.StringVar(&blogpath, "b", "", "Specify path to Org export / blog files.")
	flag.Parse()
	if blogpath == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	fmt.Println(blogpath)
}
