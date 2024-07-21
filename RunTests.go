package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Script supports only one optional argument - path to the file you want to test.")
		return
	}

	root := "."
	directories := []string{"algorithms", "structures"}

	if len(os.Args) == 2 {
		// Run a specific test file if path is provided
		testFilePath := os.Args[1]
		runTestFile(testFilePath)
	} else {
		// Traverse directories and run tests
		for _, dir := range directories {
			fullPath := filepath.Join(root, dir)
			err := filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					runTestsInDir(path)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error walking the path %q: %v\n", fullPath, err)
			}
		}
	}
}

func runTestsInDir(dir string) {

	cmd := exec.Command("go", "test", "-v")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error running tests in directory %s: %v\n", dir, err)
	}
	fmt.Printf("Output of tests in %s:\n%s\n", dir, string(output))
}

func runTestFile(testPath string) {

	separator := string(filepath.Separator)
	dirPath := testPath + separator
	dirPathFull := "." + separator + dirPath

	// Ensure the provided path is a directory
	dir, err := os.Open(dirPath)
	if err != nil {
		log.Fatalf("Error opening directory: %v", err)
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1) // Read all directory entries
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), "_test.go") && !e.IsDir() {
			fileName := filepath.Join(testPath, e.Name())

			// Run go test with the package path
			cmd := exec.Command("go", "test", "-v", dirPathFull)
			cmd.Dir = "." // Set the working directory to the module root

			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error running test file %s: %v\n", fileName, err)
				fmt.Printf("Output: %s\n", output)
				continue
			}
			fmt.Printf("Output of test file %s:\n%s\n", fileName, string(output))
		}
	}
}
