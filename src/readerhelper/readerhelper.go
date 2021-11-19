package readerhelper

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func StringFileReader(fileName string) []string {
	var content []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fmt.Errorf("error "), err)
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		content = append(content, scanner.Text()+"$$")
	}
	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return content

}
