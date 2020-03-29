package main

import (
	"fmt"
)

var sum int

func main() {
	for i := 0; i < 99; i++ {

		var input int
		fmt.Println("Please select an option:")
		fmt.Println("1 – Print Covid19 cases in Pakistan ")
		fmt.Println("2 – Print Covid19 cases in SouthKorea")
		fmt.Println("3 – Print Covid19 cases in France")
		fmt.Println("4 – Print my personalized message to Coronavirus")
		fmt.Println("0 – Exit")
		fmt.Scanln(&input)

		if input == 1 {
			pakistan()
			sum = sum + 1
		}
		if input == 2 {
			southkorea()
			sum = sum + 2
		}
		if input == 3 {
			france()
			sum = sum + 3
		}
		if input == 4 {
			msg()
		}
		if input == 0 {
			if sum >= 6 {
				i = i + 100
			}
			if sum < 6 {
				fmt.Println("you havent seen seen the information of all countries")
			}
		}

	}

}
func pakistan() {
	fmt.Println("1563 cases in Pakistan")
}

func southkorea() {
	fmt.Println("9583 cases in Pakistan")
}

func france() {
	fmt.Println("37575 cases in Pakistan")
}

func msg() {
	fmt.Println("May we get rid of this virus asap")
}
