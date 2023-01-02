package main

import (
	"bufio"
	"fmt"
	"os"
	"pomidoro/utils"
	"strings"
	"time"
)

type tomato struct {
	name     string
	duration int
}

func (t tomato) Start() {
	fmt.Printf("Time to %s\n", strings.ToLower(t.name))
	seconds := 3
	time.Sleep(time.Duration(seconds) * time.Second)
	utils.Timer(t.duration)
}

func getInput(c chan string) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	s, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c <- strings.TrimRight(s, "\n")
	close(c)
}

func getNextTomatoName(num int) string {
	var name string
	if num%2 != 0 {
		name = "Work"
	} else if num%4 == 0 {
		name = "Long rest"
	} else {
		name = "Short rest"
	}
	return name
}

func getListOfCommands(t map[string]tomato) string {
	var r strings.Builder
	for key, el := range t {
		r.WriteString(fmt.Sprintf("\"%s\" for running %s with duration of %d minutes\n", key, el.name, el.duration))
	}
	return r.String()
}

func main() {
	current := 1
	t := map[string]tomato{
		"work": {
			name:     "Work",
			duration: 25,
		},
		"sr": {
			name:     "Short rest",
			duration: 5,
		},
		"lr": {
			name:     "Long rest",
			duration: 15,
		},
	}

	for {
		fmt.Printf("Next tomato is: %s\n", getNextTomatoName(current))
		fmt.Printf("If you are ready, print \"yes\" or one of these commands:\n%s or press \"no\" to exit programm", getListOfCommands(t))
		i := make(chan string, 1)

		go getInput(i)

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
