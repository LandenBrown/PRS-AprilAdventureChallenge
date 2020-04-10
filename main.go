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
			saveExists := false
			//Find if there is an existing save
			fmt.Println("Is there an existing save?")
			//Needs to be built out, but this will check for the existence of a file
			switch saveExists {
			case true:
				fmt.Println("Load attributes and load game engine...")
			default:
				fmt.Println("No save exists, would you like to start a new game?\n1. Yes \n2. No")
				var newGameResponse string
				fmt.Scan(&newGameResponse)
				switch newGameResponse {
				case "1":
					fmt.Println("Starting New Game...")
					modules.MainGameEngine()
				default:
					fmt.Println("Exiting...")
				}
			}
		case "2":
			fmt.Println("You selected New Game")
			tutorialSelected := "0"
			//Ask if they would like to start the tutorial
			fmt.Println("If this is your first time experiencing Dark Tribal, we recommend going through the tutorial.")
			fmt.Println("Would you like to start the tutorial now?\n1. Yes\n2. No")
			fmt.Scan(&tutorialSelected)
			//Needs to be built out, but this will check for the existence of a file
			switch tutorialSelected {
			case "1":
				fmt.Println("Loading tutorial")
			default:
				fmt.Println("Skipping tutorial...")
			}
			fmt.Println("Load player option selections...")
		case "3":
			fmt.Println("You selected Options")
			fmt.Println("Needs designed by Jake...")
		case "4":
			fmt.Println("Exiting game...")
			i = 2
		default:
			fmt.Println("Invalid Option.")
		}

	}

}
