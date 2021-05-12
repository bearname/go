package main

func main() {
	logoImage := &File{"image.png"}
	architectureUml := &File{"uml.drawio"}
	presentationPowerPoint := &File{"presentation.pptx"}
	docs := &Folder{name: "img"}
	docs.add(logoImage)

	src := Folder{name: "docs"}
	src.add(architectureUml)
	src.add(presentationPowerPoint)
	src.add(docs)

	src.search("rose")
}
