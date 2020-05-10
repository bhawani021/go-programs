// Use os.File.Readdir
package main

import (
	"fmt"
	"os"
	"sort"
)

func main()  {
	dirName := "."

	f, err := os.Open(dirName)
	if err != nil {
		panic(err)
	}

	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		panic(err)
	}

	for _, file := range list {
		fmt.Println(file.Name())
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name() < list[j].Name()
	})

	fmt.Println("========================")

	for _, file := range list {
		fmt.Println(file.Name())
	}
}


