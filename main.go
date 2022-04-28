package main

import (
	"fmt"
)

func main() {
	fmt.Println("What do you want from below list")

	fmt.Println("1. List of upcomming series \n 2. List of countries playing cricket \n 3. List of Current matches \n 4. Search about your favorite player")
	var choice int

	fmt.Scan(&choice)
	var playerName string
	if choice == PlayersSearch {
		fmt.Println("Please enter your favourite player name")
		fmt.Scan(&playerName)
	}

	GetData(choice, playerName)

	fmt.Scan(&choice)
}
