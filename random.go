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
	"time"
)
/*
The main function in package main will be entry point of our executable program.
*/
func main(){
	max := 100	//maximum number for a random function
	min := 10	//minimum number for a random function
	rand.Seed(time.Now().UnixNano())	//using the seed function to
	//Seed uses the provided seed value to initialize the generator to a deterministic state.
	//Seed should not be called concurrently with any other Rand method.
	fmt.Println("Random Number is",rand.Intn(max-min))
}
