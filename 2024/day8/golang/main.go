package main

import (
	"fmt"
	"os"
)

func main() {
	file_data, _ := os.ReadFile("input.txt")
	fmt.Println(string(file_data))

	fmt.Println("Seasons greetings gamers!")
}
