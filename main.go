package main

import "fmt"

func main() {
	fmt.Println("remove me vim-go")
	for i := 1; i < 5; i++ {
		fmt.Println("remove me indent 2")
		for j := 1; j < 5; j++ {
			// fmt.Println("remove me test indent 3")
		}
	}
}

// foo is foo
func foo() {
	fmt.Println("remove me")
	return
}

func Foo() {
	return
}
