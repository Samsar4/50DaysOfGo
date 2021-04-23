package main

import (
	"fmt"
	"os"
)

func CreateFiles(defaultDir string, files ...string) {
	for _, name := range files {
		filePath := fmt.Sprintf("%s/%s.%s", defaultDir, name, "txt")

		file, err := os.Create(filePath)

		if err != nil {
			fmt.Printf("Error! file: %s: %v\n", name, err)
			os.Exit(1)
		}

		defer file.Close()

		fmt.Printf("[+] File created: %s\n", file.Name())
	}
}

func main() {

	tmp := os.TempDir()

	CreateFiles(tmp)
	CreateFiles(tmp, "testFile0")
	CreateFiles(tmp, "testFile1", "testFile2", "testFile3")
}
