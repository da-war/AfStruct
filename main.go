package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	fmt.Println(userNames)
	fmt.Println(len(userNames))
	fmt.Println(cap(userNames))

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Jane")
	userNames[0] = "Doe"
	userNames[1] = "Smith"

	fmt.Println(userNames)
	fmt.Println(len(userNames))
	fmt.Println(cap(userNames))
}
