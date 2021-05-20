package main

import (
    "fmt"
    "flag"
    "os"
)

var name string

func init(){
   flag.StringVar(&name, "name", "everyone", "greeting") 
}

func main () {
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}


	flag.Parse()
    fmt.Printf("hello , %s \n", name)
}