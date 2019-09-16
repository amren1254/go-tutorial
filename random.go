/*Go program to obtain and print random number
*/
/*
package main description
Program in Go are made up of packages.
The package main tells go compiler that package should compile as an executable program instead of shared library.
Package called main is used to create executable binary.
Program execution starts in package main by calling its function which is also called main.
*/

package main

/*
package needs to be imported first to use its exported identifiers
*/
import ("fmt"
	"math/rand"
)
/*
The main function in package main will be entry point of our executable program.
*/
func main(){
	fmt.Println("Random Number is",rand.Intn(10))
}
