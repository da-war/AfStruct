package main

import "fmt"

func main() {

	var productNames [4]string = [4]string{"p1", "p2", "p3", "p4"}
	prices := [4]float64{1.1, 2.2, 3.3, 4.4}

	fmt.Println(productNames)

	fmt.Println(prices)
	for i := 0; i < len(prices); i++ {
		fmt.Printf("Product %s costs $%.2f\n", productNames[i], prices[i])
	}

	//slice
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
}
