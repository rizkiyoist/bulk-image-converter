package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// find all filenames in this directory
	fileNames, err := findFileNames()
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("cvrt", 0777)
	for _, fileName := range fileNames {
		noExtFileName := removeExtension(fileName)
		cmd := exec.Command("magick", fileName, "-resize", "50%", "JPG:"+noExtFileName+"_resized.jpg")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(fileName, err)
		}
		fmt.Println(string(out))
		cmd = exec.Command("mv", noExtFileName+"_resized.jpg", "cvrt/"+noExtFileName+"_resized.jpg")
		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println(fileName, err)
		}
	}
	fmt.Println("done")
}

func findFileNames() ([]string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed reading files: %v", err)
	}
	fileNames := []string{}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

func removeExtension(fileName string) string {
	dot := strings.LastIndexByte(fileName, '.')
	if dot != -1 {
		fileName = fileName[:dot]
	}
	return fileName
}
