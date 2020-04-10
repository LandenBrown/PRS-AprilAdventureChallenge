package main

import (
	"fmt"

	"./modules"
)

func main() {
	versionNumber := 0.01

	fmt.Println("Welcome to Dark Tribal")
	fmt.Println("Version: ", versionNumber)

	for i := 0; i < 1; i-- {
		var startResponse string
		fmt.Println("Please select an option")
		fmt.Println("1. Load Game\n2. New Game\n3. Options\n4. Exit")
		fmt.Scan(&startResponse)
		switch startResponse {
		case "1":
			fmt.Println("You selected Load Game")
			modules.LoadGame()
		case "2":
			fmt.Println("You selected New Game")
			modules.NewGame()
		case "3":
			fmt.Println("You selected Options")
			modules.StartMenuOptions()
		case "4":
			fmt.Println("Exiting game...")
			i = 2
		default:
			fmt.Println("Invalid Option.")
		}

	}

}
