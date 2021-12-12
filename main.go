package main

import (
	"flag"
	"fmt"
)

func main() {
	var blogpath string
	flag.StringVar(&blogpath, "b", "blogpath", "Specify path to Org export / blog files.")
	flag.Parse()
	fmt.Println(blogpath)
}
