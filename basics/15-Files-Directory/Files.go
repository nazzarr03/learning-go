package main 

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create an empty file
	emptyFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emptyFile)
	emptyFile.Close()

	// Rename the file
	oldName := "test.txt"
	newName := "test2.txt"
	os.Rename(oldName, newName)

	// Check if file exists
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("File does not exist")
	}

	// Get file information
	fileStat, err := os.Stat("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileStat.Size())
	fmt.Println(fileStat.ModTime())
	fmt.Println(fileStat.Name())

	// Writing to a file
	outFile := "test2.txt"
	content := []byte("Hello world!\n")
	err = ioutil.WriteFile(outFile, content, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Writing to file successful")
	}	

	// Reading from a file
	fileName := "test2.txt"
	fileContent, fileError := ioutil.ReadFile(fileName)
	if fileError != nil {
		log.Fatalln("Error reading file:", fileError)
	} else {
		fmt.Println(string(fileContent))
	}

	// Append to an existing file
	af, err := os.OpenFile("test2.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := af.Write([]byte("Appending to existing file\n")); err != nil {
		fmt.Println(err)
	}
	af.Close()

	// Reduce the file size
	err = os.Truncate("test2.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

	// Delete a file
	err = os.Remove("test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}