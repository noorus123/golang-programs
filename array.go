package main

import (
	"fmt"
)

func main() {
	var n [10]int
	for i := 0; i < 10; i++ {
       	 n[i] = i
  	}
	n[4]=100
	fmt.Println("printing array")
	printArray(n)
		
	//convert complete array into slice	
	var slice = n[:]
	fmt.Println("printing complete slice")	
	printSlice(slice)
		
	//convert array into slice	
	var sli = n[:len(n)-1]
	fmt.Println("printing modified slice")
	printSlice(sli)
	
}
func printArray(arr [10]int){
   fmt.Printf("len=%d cap=%d Array=%v\n",len(arr),cap(arr),arr)
}

func printSlice(sli []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(sli),cap(sli),sli)
}