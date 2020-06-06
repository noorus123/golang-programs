package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "hello "	
	
	//print string
	fmt.Printf("print string : %s",str)
	
	//Access string character by character
	var i int
	for i=0;i<len(str);i++ {
		fmt.Printf("%c",str[i])
	}
	
	fmt.Printf("\n Third character of string str is %c",str[2])
		
	//concatenate two strings 
	var str2 string = "Moto"
	var str3 string
	str3 = str+str2	
	fmt.Printf("\n concatenate two strings : %s",str3)
	
	
	//combine characters to form string
	var str4 string = "chair table"
	var str5 string
    for i := 0; i < len(str4); i++ {
        str5= str5+string(str4[i])
    }	
	fmt.Printf("\n combine characters to form string : %s",str5)
	
	//Replace existing character with new character in a given string	
	var str6 string = "follow hello "
	str6 = strings.Replace(str6,"l","L",-1)
	fmt.Printf("\n Inserted characters in string : %s",str6)	
}