package main

import "fmt"

func main() {

	name := "Budi kuda lari gagah berani hehe"

	if name == "Argian Raditya" {
		fmt.Println("Hello, Argian Raditya!")
	} else if name == "argu" {
		fmt.Println("ini argu?")
	} else {
		fmt.Println("Hi, siapa kamu?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}