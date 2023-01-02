package utils

import (
	"fmt"
	"time"
)

func Timer(max int) {
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
