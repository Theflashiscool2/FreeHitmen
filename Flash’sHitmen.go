package main

import (
	"fmt"
"strings"
)

func main() {

	hm1 := "Hashim"
	hm2 := "xRyze1"
	hm3 := "B1az3z"
	var yon string
	var hmselect string

	fmt.Printf("Welcome to flashes hitmen service where we kill anyone you would like for the right price are you sure you wish to continue?\n")
	fmt.Scan(&yon)
  if !strings.EqualFold(yon, "yes") && !strings.EqualFold(yon, "no") {
			fmt.Println("That’s not a valid option it’s a yes or no question.")
    return 
    }
	if strings.EqualFold(yon, "no") {
		fmt.Println("ok see you later then. If you tell anyone about this service we will find you")
		return
	}
	 if strings.EqualFold(yon, "yes") {
		fmt.Printf("We have 3 hitmen, each have a special set of skills I will list each hitmen for you, their skill, and their price\n")
		fmt.Printf("first hitman is %v. He has the ability to hack into anything. I mean anything. He can hack into his targets toaster and burn them to fucking death somehow. His rate is 50k techits per kill\n", hm1)
		fmt.Printf("your second hitman up for hire is %v. He specializes in murdering women. His special skill is that he scares women to death with his height. All women hate short men and his height just scares them to death. His rate is 10k techits per kill.\n", hm2)
		fmt.Printf("Lastly our last hitman is %v. He specializes in murdering people with his gayness. He using his furryness and gayness to murder anyone in his path. It is best to not choose him if the person who ur murdering has any animal in their home. He has to be kept away from animals for a disclosed reason. Anyway his rate is 150k techits because he’s just too gay and too much of a furry to be cheap.\n", hm3)
		fmt.Printf("Ok now it’s time to make a selection. Which one will it be. Make sure to spell it as spelled on the list.\n")
		fmt.Scan(&hmselect)
		if strings.EqualFold(hmselect, "Hashim") {
			fmt.Printf("ok. %v will be sent to your current location to receive his money and your orders\n.", hm1)
		}
		if strings.EqualFold(hmselect, "xRyze1") {
			fmt.Printf("ok. %v will be sent to your current location to receive his money and your orders\n.", hm2)
		}
		if strings.EqualFold(hmselect, "B1az3z") {
			fmt.Printf("ok. %v will be sent to your current location to receive his money and your orders\n.", hm3)
		}
    if !strings.EqualFold(hmselect, "Hashim") && !strings.EqualFold(yon, "xRyze1") && !strings.EqualFold(yon, "B1az3z") {
      fmt.Println("Thats not a hitman we have stupid")
      return
    }
		fmt.Println("Thank you for using our service.")
		
		
	}
}

