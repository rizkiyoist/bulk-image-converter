package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	// find all filenames in this directory
	fileNames, err := findFileNames()
	if err != nil {
		fmt.Println(err)
	}

	os.Mkdir("cvrt", 0777)

	failedCount := 0
	for _, fileName := range fileNames {
		noExtFileName := removeExtension(fileName)
		cmd := exec.Command("magick", fileName, "-resize", "50%", "JPG:"+noExtFileName+"_resized.jpg")
		_, err := cmd.CombinedOutput()
		if err != nil {
			failedCount++
			fmt.Println(fileName, err)
			continue
		}
		cmd = exec.Command("mv", noExtFileName+"_resized.jpg", "cvrt/"+noExtFileName+"_resized.jpg")
		_, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("failed copying: "+fileName, err)
		}
	}
	fmt.Printf("succesfully converted %d files, failed: %d", len(fileNames)-failedCount, failedCount)
}

func findFileNames() ([]string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed reading files: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		t1, err := files[i].Info()
		if err != nil {
			fmt.Println(err)
		}
		t2, err := files[j].Info()
		if err != nil {
			fmt.Println(err)
		}
		return t1.ModTime().After(t2.ModTime())
	})

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
