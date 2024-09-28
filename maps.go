package main

import "fmt"

func maining() {
	fmt.Println("Hello, World!")
	websites := map[string]string{
		"SitePoint": "https://www.sitepoint.com",
		"Google":    "https://www.google.com",
		"Facebook":  "https://www.facebook.com",
	}

	fmt.Println(websites["SitePoint"])
	fmt.Println(websites["Google"])

	fmt.Println(websites)
	fmt.Println("sdfsafdsa", websites["Instagram"])

	websites["Instagram"] = "https://www.instagram.com"

	fmt.Println(websites)
	fmt.Println(websites["Instagram"])

	delete(websites, "SitePoint")

	fmt.Println(websites)

}
