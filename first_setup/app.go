// every go file needs a package
// you can have multiple files that make up one package
// and you can have multiple packages in one go project
// goal of packages is to organize your code
// you can use any name that you want, but using main bc main is a special package name that tells go this is the main package file/entryway of the application.
// This tells go, "THIS is where the execution should start"
package main

import "fmt"

// Double quotes or backticks ONLY for strings
// Also needs to be called main so go knows which code to execute when it starts
func main() {
	fmt.Print("Hello World!")
}

//main application code must be wrapped within the main function

// go mod init [someURL path - e.g. github or where you want to expose the code to]
// go build > creates an executable version of the file, even if someone doesnt have go installed
// then you can run ./first-app, go finds main entry point and runs the program
//if you create a change to the file and save, you need to re-run go build before executing the program to view the changes
