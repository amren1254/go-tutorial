package main

import (
	"fmt"
)
func main(){
	var num int;
	fmt.Scanf("%d",&num);
	var n,m int = 1,1;

	for n<=num {
		m=m*n;
		n=n+1;
	}
	fmt.Println(m);
}