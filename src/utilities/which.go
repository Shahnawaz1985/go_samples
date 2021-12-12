package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("custom which unix utility!!")

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide the name of file as argument for search")
		return
	}

	file := arguments[1]

	path := os.Getenv("PATH")
	fmt.Println("captured path :", path)
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			fmt.Printf("Mode : %s captured for file %s", mode, fileInfo)
			fmt.Println("-----------------------------------------------------")
			if mode.IsRegular() {
				fmt.Println("File is in regular mode!")
				if mode&0111 != 0 {
					fmt.Println("Doing positive bitwise comparison!")
					fmt.Println(fullPath)
					//return
				}
			}
		}

	}
}
