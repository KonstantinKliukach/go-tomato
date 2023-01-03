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
	current := 1
	t := tomato.GetAllTomatoes()

	for {
		fmt.Printf("Next tomato is: %s\n", getNextTomatoName(current))
		fmt.Printf("If you are ready, print \"yes\" or one of these commands:\n%s or press \"no\" to exit programm", tomato.GetListOfCommands(t))
		i := make(chan string, 1)

		go utils.GetInput(i)

		s := <-i

		if s == "yes" {
			if current%2 != 0 {
				t["work"].Start()
			} else if current%4 == 0 {
				t["sr"].Start()
			} else {
				t["lr"].Start()
			}
		} else if s == "no" {
			fmt.Println("See you, space cowboy")
			os.Exit(0)
		}
	}
}
