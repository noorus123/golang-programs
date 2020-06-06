package main

import (
	"fmt"
)
type employee struct {
		id int
		name string
		desgn string
	}	
	 
func main() {
	 var emp employee	
	  emp = employee {1,"abc","SSE"}	
	  fmt.Println("struct values",emp)
	  fmt.Println("struct values - name",emp.name)
	  fmt.Println("struct values - designation",emp.desgn)
}