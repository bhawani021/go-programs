// Use os.File.Readdir
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		*files = append(*files, path)
		return nil
	}
}

func main()  {
	root := "."

	var list []string

	err := filepath.Walk(root, visit(&list))
	if err != nil {
		panic(err)
	}

	for n, file := range list {
		fmt.Println(n, file)
	}
}


