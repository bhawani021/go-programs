package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetDir(dirName string) ([]os.FileInfo, error) {
	var files []os.FileInfo

	// ReadDir reads the directory named by dirname and returns
	// a list of directory entries sorted by filename.
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return files, err
	}

	return files, nil
}

func MakeDir(dirName string) error {
	err := os.Mkdir(dirName, 0700)

	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}


func DeleteDir(dirName string) error {
	err := os.RemoveAll(dirName)
	return err
}

func main()  {
	// Make dir
	err := MakeDir("./testDir")
	if err != nil {
		panic(err)
	}

	err = MakeDir("./testDir/testDir1")
	if err != nil {
		panic(err)
	}

	err = MakeDir("./testDir/testDir2")
	if err != nil {
		panic(err)
	}

	// Get Dir
	files, err := GetDir("./testDir")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Delete dir
	err = DeleteDir("./testDir")
	if err != nil {
		panic(err)
	}
}
