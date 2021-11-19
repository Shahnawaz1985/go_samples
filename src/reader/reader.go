package main

import (
	"bufio"
	"fmt"
	"greeter"
	"log"
	"os"
	"readerhelper"
)

func main() {
	var data [5]string
	var myslice []string
	file, err := os.Open("C:\\Intellij_Workspace\\Intellij_Idea_GoLand\\test_go_Project\\src\\reader\\test.txt")
	//file, err := os.Open("test.txt")
	//file, err := ioutil.ReadFile("C:\\Intellij_Workspace\\Intellij_Idea_GoLand\\test_go_Project\\src\\filereader\\test.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("error "), err)
		log.Fatal(err)
	}
	//fmt.Println("Contents of the file :", string(file))
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		data[i] = scanner.Text() + "--"
		i++
	}
	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Println("Filled arr :", data)
	greeter.Greet2()
	greeter.Greet1()
	fmt.Println("---------------------------------------------------------------")
	myslice = readerhelper.StringFileReader("C:\\Intellij_Workspace\\Intellij_Idea_GoLand\\test_go_Project\\src\\reader\\test2.txt")
	fmt.Println("Data :", myslice)
	fmt.Println("Length of slice : ", len(myslice))

}
