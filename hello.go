/*First hello world program to print hello world  */
/*
Programs in Go are made up of packages.
The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.
The main function in the package “main” will be the entry point of our executable program.
Package called main is used to create executable binary.
Program execution starts in package main by calling its function which also called main.
*/
package main

/*
Package needs to be imported first to use its exported identifiers. It’s done with construct called import declaration
Below we’ve one import declaration with one ImportSpec. Each ImportSpec defines single package import.
*/
import "fmt"	//fmt(format) is package to format number and strings

func main(){

	fmt.Println("hello World")	//Println will print the string or number with new line.
}
