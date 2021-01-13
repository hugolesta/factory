package main

import (
	"fmt"
	"os"
)

func main() {
	var t int
	fmt.Printf("What's the desired connection? ")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error reading the option: %v", err)
		os.Exit(1)
	}
	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Engine is not valid")
		os.Exit(1)
	}
	t, err := conn.GetNow()

	if err != nil {
		fmt.Printf("It couldn't get the date: %v", err)
		os.Exit(1)
	}
	fmt.Println(now.Format("2020-01-13"))
	err := conn.Close()
	if err != nil {
		fmt.Println("It coldn't close the connextion %v", err)
	}
}
