/* Go program to use math package */
/*
Every Go program is made up of package.
The package main tells that package should compile as an executable program instead of shared library
*/

package main

import ("fmt"
	"math")

func main(){
	fmt.Println("You have %g Problems",math.Sqrt(7))
}
