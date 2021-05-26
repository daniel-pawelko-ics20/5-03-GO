// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that returns what movie a user can watch alone

package main

import (
	"fmt"
)

// This main function will ask user for age and return what the user can watch
func main() {
	// Defining variables
	var age int

	fmt.Println("What movie can you watch?")
	fmt.Println()

	// User Input
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	fmt.Println()

	// Return what user can watch
  switch{
    case age >= 17:
		  fmt.Println("You can see an R movie alone.")
    case age >= 13:
      fmt.Println("You can see a PG-13 movie alone.")
    case age >= 5:
      fmt.Println("You can see a G or PG movie alone.")
    default:
      fmt.Println("Uh. You're too young for most things.")
  }
}
