package main

import (
	"fmt"
)

func main() {

	myCourses := []string{"Docker", "Puppet", "Gayboy"}

	fmt.Printf("Length is: %d.\nCapacity is: %d ", len(myCourses), cap(myCourses))
	//map[keyType]valueType

}
