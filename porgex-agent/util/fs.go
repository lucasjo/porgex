package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExistsDir(path string) bool {

	if path == "" {
		return false
	}

	dir, err := os.Stat(path)
	ErrorCheck(err)

	if !dir.IsDir() {
		fmt.Errorf("Could is not a Directory")
		return false
	}

	return true

}

func Searchdirlist(searchDir string) {

	isdir := ExistsDir(searchDir)

	fmt.Printf("is dir : %v\n", isdir)

	filelist := []string{}

	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, e error) error {

		filelist = append(filelist, f.Name())

		return nil
	})

	ErrorCheck(err)

	for _, file := range filelist {
		fmt.Printf("%v\n", file)
	}

}
