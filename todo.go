package main

import (
	"bufio"
	"fmt"
	"os"
)

// By default, show an empty to-do tick box and the text "Item 1" as
// backdrop message. Then once the user types, that becomes the actual
// to-do list item. The user can add more boxes by pressing enter.
// The user can toggle the tick boxes by clicking on the box.

func main() {
	var todoList []string
	fmt.Println("To-do:")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Add another task or 'q': ")

		if !scanner.Scan() {
            break
        }

		line := scanner.Text()
		
		if (line == "q") {
			break;
		}
		
		todoList = append(todoList, line)
        fmt.Println("\nCurrent To-do List:")

		for _, item := range todoList {
			fmt.Printf("[ ] %s\n", item)
		}
		fmt.Println("--------------------------")

	}
}