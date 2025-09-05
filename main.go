package main

import "fmt"

func main() {
	var (
		customerName string = "Argian Raditya"
		totalCoin int = 101
		coinPrice int = 3500 // in IDR
		values = [3]int{
			10, 
			20, 
			30,
			}	
	)

	if totalCoin > 100 {
		totalCoin += 10 // bonus 10 coins
	}

	fmt.Println("Hi ", customerName, ", our beloved customer")
	fmt.Println("Congratulations! You have received 10 bonus coins")
	fmt.Println("You have ", totalCoin, " coins")
	fmt.Println("Each coin is worth $", coinPrice)
	fmt.Println("So your total balance is $", totalCoin * coinPrice)
	fmt.Println("Your values are ", values[2])
	fmt.Println("len(values) = ", len(values) )
}
