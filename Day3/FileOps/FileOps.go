package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(origPath, newPath string) {
	origFile, err := os.Open(origPath)
	check(err)
	defer origFile.Close()

	// Create destination file
	newFile, err := os.Create(newPath)
	check(err)
	defer newFile.Close()

	// Copy bytes from original to new file
	bytesWritten, err := io.Copy(newFile, origFile)
	check(err)
	log.Printf("Copied %d bytes", bytesWritten)

	// Commit
	err = newFile.Sync()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RenameFile(origFilename, newFilename string) {
	err := os.Rename(origFilename, newFilename)
	check(err)
}

func GetFileInfo(filename string) {
	fileInfo, err := os.Stat(filename)
	check(err)
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %v\n", fileInfo.Sys())
}

func TruncateFile(filename string, size int64) {
	err := os.Truncate(filename, size)
	check(err)
}

func CreateEmptyFile(filename string) {
	newFile, err := os.Create(filename)
	check(err)
	log.Println(newFile)
	newFile.Close()
}

func main() {
	//CreateEmptyFile("testfile.txt")
	//TruncateFile("testfile.txt", 35)
	//GetFileInfo("testfile.txt")
	//RenameFile("/tmp/testfile.txt", "testfile.txt")
	CopyFile("testfile.txt", "/tmp/testfile.txt")
}
