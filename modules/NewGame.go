package modules

import (
	"fmt"

	"./xtra"
)

func NewGame() {
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
	xtra.TestXtra()
}
