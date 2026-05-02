package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			fmt.Print("invalid string")
			return
		}
		cInput := cleanInput(input)
		firstWord := cInput[0]
		if cmd, ok := commandMap[firstWord]; ok {
			cmd.callback()
		} else {
			fmt.Println("unknown command:", firstWord)
		}
	}
}
