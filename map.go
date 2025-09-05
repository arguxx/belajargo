package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Argian Raditya",
		"address": "Jakarta",
		"phone":   "08123456789",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("Address:", person["address"])

	person["title"] = "Programmer"
	fmt.Println("Title:", person["title"])

	delete(person, "address")
	fmt.Println("After deletion, address:", person["address"])

	fmt.Println("Complete map:", person)
	fmt.Println("Length of map:", len(person))
}