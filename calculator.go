package main

import (
	"fmt"
)
func calculator(arg1,arg2 int, operator string) (result int){
	switch(operator){
	case "add": result=arg1+arg2
	case "sub": result=arg1-arg2
	case "mul": result=arg1*arg2
	case "div": result=arg1/arg2
	}
	return
}

func main() {
	a := 1
	b := 2
	fmt.Println(calculator(a,b,"add"))
}