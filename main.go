package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	// Check Args
	if len(os.Args) < 2 {
		panic("not enough arguments (specify project name)")
	}

	if os.Args[1] == "help" {
		help()
		return
	}

	// create C# sln
	createSln()

	if len(os.Args) < 3 {
		createGitignore()
		createReadme(os.Args[1])
		return
	}

	// check Flags
	switch os.Args[2] {
	case "-ng":
		fmt.Println("No .gitignore file will be created")
		createReadme(os.Args[1])
	case "-nr":
		fmt.Println("No README file will be created")
		createGitignore()
	case "-ngr":
		fmt.Println("No .gitignore and README.md files will be created")
	case "-nrg":
		fmt.Println("No .gitignore and README.md files will be created")
	default:
		createGitignore()
		createReadme(os.Args[1])
	}

}
