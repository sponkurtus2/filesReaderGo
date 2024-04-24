package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Custom type for our files
type FileInfo map[string]string

// Func to check for errors
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (f FileInfo) Add(file, extension string) {
	f[file] = extension
}

func main() {
	// Location which we want to check
	location, err := os.ReadDir(".")
	check(err)

	// Get the current files from the directory
	// files := []string{}
	// for _, file := range location {
	// 	files = append(files, file.Name())
	// }

	// // Get the extension of the previous files
	// filesType := []string{}

	// for _, file := range files {
	// 	fileExtension := path.Ext(file)
	// 	filesType = append(filesType, fileExtension)
	// }

	finalFiles := make(FileInfo)

	for _, file := range location {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		fileExtension := filepath.Ext(fileName)
		finalFiles.Add(fileName, fileExtension)

	}

	// for i := 0; i < len(files); i++ {
	// 	file := files[i]
	// 	extension := filesType[i]
	// 	finalFiles.Add(file, extension)
	// }

	// for i := range files {
	// 	finalFiles.Add(files[i], filesType[i])
	// }

	for file, extension := range finalFiles {
		fmt.Printf("%q, %q \n", file, extension)
	}

}
