package engine

/*
TO DO 4/11/2020:
Create functions for Camp Huts that makes it cost more to upgrade


*/

import (
	"fmt"

	"./gameobjects"
)

func UpgradeBuilding(x *gameobjects.Hut, y *gameobjects.Player) {
	plusOneLevel := x.Level + 1
	fmt.Println("The", x.Name, "upgrade from level", x.Level, "to level", plusOneLevel,
		"will cost ", x.WoodReq, "Wood,", x.StoneReq, "Stone, and",
		x.FiberReq, "Fiber.")
	fmt.Println("\nContinue?\n1. Yes\n2. No")
	upgradeChoice := ""
	fmt.Scan(&upgradeChoice)
	if upgradeChoice == "1" {
		if y.Wood >= x.WoodReq {
			if y.Stone >= x.StoneReq {
				if y.Fiber >= x.FiberReq {
					x.Level++
					fmt.Println("Your", x.Name, "is now level", x.Level)
				} else {
					fmt.Println("You don't have enough fiber!")
				}
			} else {
				fmt.Println("You don't have enough stone!")
			}
		} else {
			fmt.Println("You don't have enough wood!")
		}
	}
}

func ManageCampMenu() {
	i := 0
	for i <= 2 {
		fmt.Println("\n\n-------Camp:", ThePlayer.Camp.Name, "-------")
		fmt.Println("1. Operations Hut: Level", ThePlayer.Camp.OperationsHut.Level)
		fmt.Println("2. Supplies Hut: Level", ThePlayer.Camp.SuppliesHut.Level)
		fmt.Println("3. Villager's Hut: Level", ThePlayer.Camp.VillagersHut.Level)
		fmt.Println("4. Back")
		fmt.Println("-------------------------")
		campMenu := ""
		fmt.Println("Which building would you like to inspect?")
		fmt.Scan(&campMenu)
		switch campMenu {
		case "1":
			UpgradeBuilding(&ThePlayer.Camp.OperationsHut, &ThePlayer)
		case "2":
			UpgradeBuilding(&ThePlayer.Camp.SuppliesHut, &ThePlayer)
		case "3":
			UpgradeBuilding(&ThePlayer.Camp.VillagersHut, &ThePlayer)
		case "4":
			i = 3
		default:
			fmt.Println("Invalid Response")
		}
	}

}
