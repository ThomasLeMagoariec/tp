package main

import (
	"fmt"
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createGitignore() {
	// Create .gitignore file
	f, err := os.Create(".gitignore")
	check(err)
	defer f.Close()

	_, err = f.WriteString("bin/\nobj/\n\n.idea/\n*~\n*.DotSettings.user\n")
	check(err)

	gac("created .gitignore")
}

func createReadme(name string) {
	// Create README.md file
	f, err := os.Create("README")
	check(err)
	defer f.Close()

	_, err = f.WriteString(name + "\n\n")
	check(err)

	gac("created README")
}

func createSln() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Creating project %s in %s\n", os.Args[1], dir)

	cmd := exec.Command("dotnet", "new", "sln", "--name", os.Args[1])
	cmd.Dir = dir

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Print(string(stdout))

	cmd2 := exec.Command("dotnet", "new", "console", "-n", os.Args[1], "-f", "net7.0", "-lang", "C#")
	cmd2.Dir = dir

	stdout2, err2 := cmd2.Output()

	fmt.Print(string(stdout2))
	if err2 != nil {
		panic(err2)
	}

	cmd3 := exec.Command("dotnet", "sln", "add", os.Args[1]+"/"+os.Args[1]+".csproj")
	cmd3.Dir = dir

	stdout3, err3 := cmd3.Output()

	if err3 != nil {
		panic(err3)
	}

	fmt.Print(string(stdout3))

	gac("created solution")
}

func gac(commit string) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "add", ".")
	cmd.Dir = dir

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Print(string(stdout))

	cmd2 := exec.Command("git", "commit", "-m", commit)
	cmd2.Dir = dir

	stdout2, err2 := cmd2.Output()

	if err2 != nil {
		panic(err2)
	}

	fmt.Print(string(stdout2))
}

func help() {
	fmt.Println("TP v1.0.0 | Created with <3 by Thomas")

	fmt.Println("USAGE:")
	fmt.Println("\ttp <project name> [FLAGS]")

	fmt.Println("FLAGS:")
	fmt.Println("\t-ng:\twon't create a .gitignore file")
	fmt.Println("\t-nr:\twon't create a README file")
	fmt.Println("\t-ngr:")
	fmt.Println("\t-nrg:\twon't create a .gitignore or a README file")
}
