package main

import "fmt"

type specialString string

var whatIsThis specialString = "hunzi"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case int:
			fmt.Printf("any %v is a int type", v)
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}
