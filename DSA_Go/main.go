package main
//package binary

import (
	"fmt"
	//"binary"
)

func main(){
	fmt.Println("Hello Goland");
	var array = make([]int,10);
	//array in ascending order
	array[0] = 5;
	array[1] = 10;
	array[2] = 19;
	array[3] = 25;array[7] = 79;
	array[4] = 35;
	array[5] = 54;
	array[6] = 65;
	array[8] = 83;array[9] = 99;
	for i:=0;i<10;i++{
		fmt.Println(array[i]);
	}
	result,searchcount := binary.binarySearch(array, 99);
	fmt.Println(result,searchcount);
	//fmt.Println(dsa_go);
}