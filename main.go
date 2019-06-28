package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var availableLocations = []string{
	"aluu", "choba", "alakahia",
	"rumuosi", "rumuokoro", "mgbuoba",
	"rumuola",
}

func getPickUpLocationFromUser() string {

	text := getAns("\nWhere are you at? ")

	for _, v := range availableLocations {
		if v == text {
			return v
		}
	}

	fmt.Println("\nThis pickup location is not available Yet?")

	return getPickUpLocationFromUser()

}

func getDropOffLocationFromUser(pickUp string) string {

	text := getAns("\nWhere are you going?")

	if text == pickUp {

		fmt.Println("\nMy Friend! Choose a different location!!")

		return getDropOffLocationFromUser(pickUp)
	}

	for _, v := range availableLocations {
		if v == text {
			return v
		}
	}

	fmt.Println("\n We don't go here yet?; choose from available locations!")

	return getDropOffLocationFromUser(pickUp)

}

func getAns(msg string) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg + " ")
	text, _ := reader.ReadString('\n')

	return strings.ToLower(strings.TrimSpace(text))
}

func getPayment(tFare int, i int) int {

	i++

	text := getAns("\nEnter Payment tFare? ")

	input, err := strconv.Atoi(text)
	if err != nil {

		fmt.Println("\nType in Valid Numbers!!")
	}

	if i > 10 {

		fmt.Println("\nEnjoy Jail time, Bye!!")

	} else if i >= 5 {

		fmt.Println(" \nDon't Be Unfortunate! I'm calling the Cops!!")
	}

	if input == tFare {

		fmt.Println("\nAwesome!")

		return input

	} else if input > tFare {

		change := input - tFare

		fmt.Printf("\nHere's Your change Sir: ₦%d ", change)

		return input

	} else if input < tFare {

		fmt.Println("\nNice try smarty pants! Pay in full Sir!")

	}
	return getPayment(tFare, i)
}

func getTip(tFare int) int {
	text := getAns("\nPlease  tip me Sir (enter tip): ")

	tip, err := strconv.Atoi(text)
	if err != nil {

		fmt.Println("\nType is valid numbers!!")

		return getTip(tFare)
	}

	if tip < 1 {
		fmt.Println("\nYou are Stingy fool!!")

	} else {
		fmt.Println("Gracias mucho!")
	}
	return tip
}

func main() {

	transport := Transport{locations: availableLocations, tFare: 50}

	fmt.Printf(" \nAvailable Locations: %s. \n\n", strings.Join(availableLocations, ","))

	pickUp := getPickUpLocationFromUser()

	transport.setPickUpLocation(pickUp)

	dropOff := getDropOffLocationFromUser(transport.pickUp)

	transport.setDropOffLocation(dropOff)

	transport.setTFare()

	fmt.Printf(" \nTP from %s to %s is: ₦%d. \n\n", transport.pickUp, transport.dropOff, transport.tFare)

	fmt.Println("Driving...")

	time.Sleep(2 * time.Second)

	fmt.Printf(" \nWe got here. Pay Up ₦%d.", transport.tFare)

	getPayment(transport.tFare, 0)
	// fmt.Println(availableLocations[0])

	tip := getTip(transport.tFare)

	transport.setTip(tip)

}
