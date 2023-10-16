package filesystem

import (
	"fmt"
	"strings"
)

type Folder struct {
	Name string
	Parent *Folder
	Children []*Folder
}

type Filesystem struct{
	Root *Folder
}

func CreateFS() *Filesystem{
	rootFolder := &Folder{
		Name: "/",
	}

	return &Filesystem{
		Root: rootFolder,
	}
}

func (folder *Folder) CreateFolder(name string) *Folder{
	newFolder := &Folder{
		Name: name,
	}
	newFolder.Parent = folder
	folder.Children = append(folder.Children, newFolder)
	return newFolder
}

func PrintFS(folder *Folder, level int) {
	queue := []Folder{}
	queue = append(queue, *folder)

	for len(queue) != 0 {
		fmt.Print(queue[0].Name)
		children := queue[0].Children
		queue = queue[1:]
		for _, child := range children{
			queue = append(queue, *child)
			fmt.Println()
			space := strings.Repeat(" ", level)
			fmt.Print(space)
			fmt.Print("|_")
			PrintFS(&queue[0], level+1)
			queue = queue[1:]
		}
	}
}

