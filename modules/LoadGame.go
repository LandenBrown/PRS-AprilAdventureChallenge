package modules

import (
	"fmt"
)

func LoadGame() {
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
		default:
			fmt.Println("Exiting...")
		}
	}

}
