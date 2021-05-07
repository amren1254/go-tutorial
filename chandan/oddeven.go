package main 

import (
	"fmt"
)

func main(){
	//write a program to check a given no is odd or even.
	var number2 int;
	fmt.Scanf("%d",&number2);
	if(number2%2==0){
		fmt.Println("Even")
	}else{
		fmt.Printf("%s","Odd");
	}
}