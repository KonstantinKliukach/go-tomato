package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func naiveTimer(max int) {
	ticker := time.NewTicker(1 * time.Second)
	isDone := make(chan bool)

	whenOver := time.Now().Add(time.Duration(max) * time.Minute)
	go func() {
		for {
			select {
			case <-isDone:
				return
			case t := <-ticker.C:
				diff := whenOver.Sub(t)
				out := time.Time{}.Add(diff)
				fmt.Print("\033[H\033[2J")
				fmt.Println(out.Format("15:04:05"))
			}
		}
	}()
	time.Sleep(time.Duration(max) * time.Minute)
	ticker.Stop()
	isDone <- true
}

type tomato struct {
	name     string
	duration int
}

func (t tomato) Start() {
	fmt.Printf("Time to %s\n", strings.ToLower(t.name))
	seconds := 3
	time.Sleep(time.Duration(seconds) * time.Second)
	naiveTimer(t.duration)
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
	fmt.Println(t)
	for {
		fmt.Printf("Next tomato is: %s\n", getNextTomatoName(current))
		fmt.Println("If you are ready, print \"yes\"")
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
