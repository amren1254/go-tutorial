package main

import(
	"fmt"
)

func binarySearch(a []int, search int) (result int, searchCount int) {
    mid := len(a) / 2
    switch {
    case len(a) == 0:
        result = -1 // not found
    case a[mid] > search:
        result, searchCount = binarySearch(a[:mid], search)
    case a[mid] < search:
        result, searchCount = binarySearch(a[mid+1:], search)
        if result >= 0 { // if anything but the -1 "not found" result
            result += mid + 1
        }
    default: // a[mid] == search
        result = mid // found
    }
    searchCount++
    return
}
func main(){
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
	result,searchcount := binarySearch(array, 99);
	fmt.Println(result,searchcount);
}