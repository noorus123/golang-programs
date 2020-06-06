var variableName map[key_data_type]value_data_type = make(map[key_data_type]value_data_type)

package main

import (
	"fmt"
)

func main() {
	var m map[string]string = make(map[string]string)
	m["up"] = "lko"
	m["mp"] = "bhopal"
	m["ap"] = "hyd"
	for state := range m {
		fmt.Printf("\ncapital of %s is %s ", state, m[state])
	}	
	city, isPresent := m["up"]  // this is boolean notation of GOlang, if the value is true then assigned to given variable(here, city) otherwise given variable(here, city) value is empty; also other variable(here, isPresent) holds true/false boolean vale
	if(isPresent){
		fmt.Printf("\ncity is there - %t and name is %s ", isPresent,city )
	}else{
		fmt.Printf("\ncity is there - %t and name is %s ", isPresent,city )
	}
}
