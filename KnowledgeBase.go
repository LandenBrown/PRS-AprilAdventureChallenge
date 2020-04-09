//This statement tells Go that this is the main package, and is to be ran, rather than supply data to another main package
package main

//This statement says that we are importing packages/libraries (or libs) for their contained functions.
import (
	"fmt"
)

//This statement is a special function in that it tells Go to call it and run it immediately
func main() {

	//This allows us to print out a *new* line to the console.
	fmt.Println("I am printing a line! Hello World!")
	fmt.Println("How are you today?")

	var myResponse string //This is how we declare variables in Go. We are going to use this variable to store the user's response.

	fmt.Scan(&myResponse) //The Scan() function exists in the "fmt" package and allows us to scan and wait for user input.
	//Above, we are telling the computer that myResponse is going to be whatever value I respond with.
	//Remember that we place the '&' symbol in front of the variable to tell the computer the memory address of the variable
	//so it can change it
	fmt.Println("You responded with:", myResponse)
	//By adding a comma and placing the variable, we are adding the value of that variable to our print function

	//Below is a
	myCourses := []string{"Docker", "Puppet", "Gayboy"}

	fmt.Printf("Length is: %d.\nCapacity is: %d ", len(myCourses), cap(myCourses))

	//map[keyType]valueType
	/*
		Testing and keeping map syntax for use in the future.
		leagueTitles := make(map[string]int)
		leagueTitles["Sunderland"] = 6
		leagueTitles["Newcastle"] = 4

	*/

}
