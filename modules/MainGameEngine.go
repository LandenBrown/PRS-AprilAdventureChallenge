package modules

import (
	"fmt"

	"./engine"
)

var gameYear = 1615
var gameMonth = 1

func EndTurn() {
	engine.ThePlayer.TurnsLeft = engine.ThePlayer.TotalTurns
	gameMonth++
	if gameMonth >= 13 {
		gameYear++
		gameMonth = 1
	}
}
func MainGameEngine() {
	engineResponse := ""
	//Starting game engine loop
	for engineResponse != "7" {
		//Test turn by turn attributes
		fmt.Println("\n----------Status---------- ")
		fmt.Println("Camp:", engine.ThePlayer.Camp.Name, "// Population:", engine.ThePlayer.Camp.Population)
		fmt.Println("Current Objective: X")
		fmt.Println("Year:", gameYear, "//", "Month:", gameMonth)
		fmt.Println("You have", engine.ThePlayer.TurnsLeft, "turns left")
		fmt.Println("--------------------------\n ")
		fmt.Println("Please select an option")
		fmt.Print("1. Manage Camp\n2. Inspect Supplies\n3. Socialize\n4. Operations\n5. Next Month\n6. Save\n7. Exit to main menu\n")
		fmt.Scan(&engineResponse)

		switch engineResponse {
		case "1":
			engine.ManageCampMenu()
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
			EndTurn()
			//engine.NextMonth()
		case "6":
			fmt.Println("Selected Save")
			//engine.Save()
		case "7":
			fmt.Println("Selected Exit")
			engineResponse = "7"
			//engine.Exit()
		default:
			fmt.Println("Invalid Response.")
		}

	}
}
