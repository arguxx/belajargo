package main

import "fmt"

func main() {
	// for with statement
	for i := 0; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
  // native for
	kuda := 0
	for kuda < 5 {
		fmt.Println("Kuda lari gagah berani", kuda)
		kuda++
	}
}