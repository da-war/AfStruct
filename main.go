package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	fmt.Println(userNames)
	fmt.Println(len(userNames))
	fmt.Println(cap(userNames))

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)

	courseRating := make(map[string]float64, 4)

	courseRating["Java"] = 4.2
	courseRating["Go"] = 4.7
	courseRating["Python"] = 4.8

	fmt.Println(courseRating)
	fmt.Println(len(courseRating))
}
