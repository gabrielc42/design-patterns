package main

import "fmt"

func main() {
	file1 := &file{name: "File Uno"}
	file2 := &file{name: "File Dos"}
	file3 := &file{name: "File Tres"}
	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder Uno",
	}
	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder Dos",
	}
	fmt.Println("\n jerarquía de impresión para la carpeta dos")
	folder2.print(" ")
	cloneFolder := folder2.clone()
	fmt.Println("\n jerarquía de impresión para la carpeta de clones")
	cloneFolder.print(" ")
}
