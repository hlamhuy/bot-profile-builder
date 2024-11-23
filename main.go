package main

import (
	"fmt"
	"profile-builder/bots"
)

func main() {
	var bot string
	fmt.Print("PROFILE CONVERTER\n1.Cyber\n2.Stellar\nChoose a bot: ")
	fmt.Scan(&bot)
	for {
		switch bot {
		case "1":
			bots.Cyber()
			return
		case "2":
			bots.Stellar()
			return
		default:
			fmt.Println("Invalid input!\n\n1.Cyber\n2.Stellar\nChoose a bot:")
			fmt.Scan(&bot)
		}
	}

}
