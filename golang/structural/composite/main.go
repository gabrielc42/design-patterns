// composite:
// treating individual objects and compositions of objects in the same way

package main

func main() {
	file1 := &file{name: "design-patterns.md"}
	file2 := &file{name: "mega-super-important-project-copy.txt"}
	file3 := &file{name: "funny-cat.jpg"}
	folder1 := &folder{
		name: "Important",
	}
	folder1.add(file3)
	folder2 := &folder{
		name: "Stuff",
	}
	folder2.add(file1)
	folder2.add(file2)
	folder2.add(folder1)
	folder2.search("mega")
}
