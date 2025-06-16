package main

import (
	"fmt"
)

func getName() string {
	name := ""

	fmt.Println("SlotMac: Slot Machine")
	fmt.Printf("Enter the user name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %v, Game Starts Now !!\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet (Balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet can't be larger than balance")
		} else {
			break
		}
	}
	return bet
}

func main() {
	balance := uint(2000)
	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
	}
	fmt.Printf("you're left with, $%d\n", balance)
}
