// This program will calculate how much time is needed to heat food.
// Created By: Marlon Poddalgoda
// Created on: 2021-04-26

package main

import "fmt"

// constants
const noPercIncrease, fiftyPercIncrease, hundredPercIncrease float64 = 1, 1.5, 2
const subTime, pizzaTime, soupTime float64 = 1, 0.75, 1.75
const sixtySeconds float64 = 60

func main() {
	// this function takes in user input and calculates cook time

	fmt.Println("This program calculates the total cook time for certain foods.")
	fmt.Println("Enter the food item you are heating.")
	fmt.Println("")

	// variables
	var foodItem string
	var cookTime float64
	var cookSeconds float64
	var numOfItems string

	// for loop
	for {
		// Ask for user input for food item
		fmt.Print("Valid choices are -> sub, pizza, soup: ")

		fmt.Scanln(&foodItem)

		// check string value
		if foodItem == "sub" {
			// assigning cook time for sub
			cookTime = subTime
		} else if foodItem == "pizza" {
			// assigning cook time for pizza
			cookTime = pizzaTime
		} else if foodItem == "soup" {
			// assigning cook time for soup
			cookTime = soupTime
		} else {
			// error catch
			fmt.Println("That's not an option! Try again.")
			continue
		}

		// number of items input
		for {
			// ask for user input for number of items
			fmt.Print("Enter how many " + foodItem + "(s) you are heating (max. 3): ")

    		fmt.Scanln(&numOfItems)

    		// process
    		if numOfItems == "1" {
    			// cooktime for 1 item
    			cookTime = cookTime * noPercIncrease
    		} else if numOfItems == "2" {
    			// cooktime for 2 items
    			cookTime = cookTime * fiftyPercIncrease
    		} else if numOfItems == "3" {
    			// cooktime for 3 items
    			cookTime = cookTime * hundredPercIncrease
    		} else {
    			fmt.Println("That's not an option, try again.")
    			continue
    		}
    		cookSeconds = cookTime * sixtySeconds
    		break
		}
		// output
		fmt.Println("")
		fmt.Printf("The total cook time is %v minutes (%v seconds).", cookTime,
					cookSeconds)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Done.")
		break
	}
}
