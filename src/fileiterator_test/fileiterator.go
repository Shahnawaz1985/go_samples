package fileiterator_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func ReportPanic() {
	e := recover()

	if e == nil {
		return
	}
	tErr, ok := e.(error)
	if ok {
		fmt.Println("Captured error :", tErr)
	}
}

func ScanDirectory(path string) {
	fmt.Println("path to traverse :", path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
