// Use os.File.Readdir
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {
	root := "."

	var list []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		list = append(list, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for i, file := range list {
		fmt.Println(i, file)
	}
}


