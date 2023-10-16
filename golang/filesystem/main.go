package main

import (
	"filesystem/filesystem"
	"fmt"
)

func main(){
	fmt.Println("Creating folder")
	fs := filesystem.CreateFS()
	root := fs.Root
	root.CreateFolder("folder 1")
	root.CreateFolder("folder 2")
	folder3 := root.CreateFolder("folder 3")
	folder31 := folder3.CreateFolder("folder 31")
	folder31.CreateFolder("folder 311")
	root.CreateFolder("folder 4")
	filesystem.PrintFS(root, 1)
}