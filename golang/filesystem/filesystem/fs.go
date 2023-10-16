package filesystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fs() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Unknown input")
			continue
		}
		text = strings.Replace(text, "\n", "",-1)

		switch text {
		case "hi":
			fmt.Println("Hello")
		case "ll":
			fmt.Println("list")
		case "../":
			fmt.Println("go back")
		case "exit":
			return
		default:
			fmt.Println("Unkown command")
		}
	}
	
}