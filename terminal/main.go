package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println("Running this application...")
	time.Sleep(1 * time.Second)
	clearTerminal()
	fmt.Println("Hello World!")
	fmt.Println("Welcome to snowball chat :)")
	fmt.Println("Do you want to join our chat room?")
JoinRoom:
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("y) Yes")
		fmt.Println("n) No")
		scanner.Scan()
		answer := scanner.Text()
		fmt.Printf("%s your answer\n", answer)
		switch answer {
		case "y":
			fmt.Println("Welcome to join the room!")
			break JoinRoom
		case "n":
			fmt.Println("See you again!")
			break JoinRoom
		default:
			clearTerminal()
			fmt.Println("Please choose one of two options.")
			continue
		}
	}
}
