package modules

import (
	"fmt"

	"./engine"
)

func MainGameEngine() {

	engineResponse := ""
	//Starting game engine loop
	for engineResponse != "7" {
		//Test turn by turn attributes
		fmt.Println("\n----------Status---------- ")
		fmt.Println("Population: X")
		fmt.Println("Current Objective: X")
		fmt.Println("Year: 1615 // Month: 1")
		fmt.Println("You have X turns left")
		fmt.Println("--------------------------\n ")
		fmt.Println("Please select an option")
		fmt.Print("1. Manage Camp\n2. Inspect Supplies\n3. Socialize\n4. Operations\n5. Next Month\n6. Save\n7. Exit to main menu\n")
		fmt.Scan(&engineResponse)

		switch engineResponse {
		case "1":
			fmt.Println("Selected Manage Camp")
			//engine.ManageCamp()
		case "2":
			fmt.Println("Selected Inspect Supplies")
			//engine.InspectSupplies()
		case "3":
			fmt.Println("Selected Socialize")
			//engine.Socialize()
		case "4": 
			fmt.Println("Selected Operations")
			//engine.Operations()
		case "5":
			fmt.Println("Selected Next Month")
			//engine.NextMonth()
		case "6":
			fmt.Println("Selected Save")
			//engine.Save()
		case "7":
			fmt.Println("Selected Exit")
			engineResponse = "7"
			//engine.Exit()

		}

	}
	fmt.Println("This will be the main game engine")
	engine.TestXtra()
}
