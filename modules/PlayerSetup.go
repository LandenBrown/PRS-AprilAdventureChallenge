package modules

import (
	"fmt"

	"./engine"
)

func PlayerSetup() {
	fmt.Println("Starting player set up...")
	fmt.Println("What is your name?")
	fmt.Scan(&engine.ThePlayer.Name)
	fmt.Println("Your name is now", engine.ThePlayer.Name)
	fmt.Println("What is your camp name?")
	fmt.Scan(&engine.ThePlayer.Camp.Name)
	fmt.Println("Your camp name is now", engine.ThePlayer.Camp.Name)
}
