/* Go program to use math package */
/*
Every Go program is made up of package.
The package main tells that package should compile as an executable program instead of shared library
*/

package main
//using the import identifier
import ("fmt"
	"math")		//using the math package to perform operation	

func main(){
	//the below function is not working.
	fmt.Println("You have %g Problems",math.Sqrt(7))
}
