package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(c chan string) {
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
