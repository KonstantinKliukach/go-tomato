package main

import (
	"fmt"
	"os"
	"pomidoro/tomato"
	"pomidoro/utils"
)

func getNextTomatoName(num int) string {
	var Name string
	if num%2 != 0 {
		Name = "Work"
	} else if num%4 == 0 {
		Name = "Long rest"
	} else {
		Name = "Short rest"
	}
	return Name
}

func main() {
	// current := 1
	allT := tomato.GetAllTomatoes()

	for {
		// fmt.Printf("Next tomato is: %s\n", getNextTomatoName(current))
		fmt.Printf("If you are ready, print \"yes\" or one of these commands:\n%s or press \"no\" to exit programm", tomato.GetListOfCommands(allT))
		i := make(chan string, 1)

		go utils.GetInput(i)

		s := <-i

		if s == "no" {
			fmt.Println("See you, space cowboy")
			os.Exit(0)
		}

		value, ok := allT[s]

		if ok {
			value.Start()
		}
	}
}
