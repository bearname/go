package main

import (
	"fmt"
)

type Folder struct {
	components []Component
	name       string
}

func (folder *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, folder.name)
	for _, composite := range folder.components {
		composite.search(keyword)
	}
}

func (folder *Folder) add(component Component) {
	folder.components = append(folder.components, component)
}
