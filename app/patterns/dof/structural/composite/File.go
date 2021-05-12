package main

import "fmt"

type File struct {
	name string
}

func (file *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, file.name)
}
