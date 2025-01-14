package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Docker Image for Golang CLI App")

	os.Mkdir("folder1", 0755)
	os.Mkdir("folder2", 0755)
	os.Mkdir("folder3", 0755)
	os.MkdirAll("multi-folder/s1/s2/s3", 0755)
}


