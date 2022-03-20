package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kiibo382/grpc_playground/service"
)

func main() {
	fmt.Println("start Rock-paper-scissors game.")
	scanner := bufio.NewScanner(os.Stdin)
	is_finished := false

	for !is_finished {
		fmt.Println("1: play game")
		fmt.Println("2: show match results")
		fmt.Println("3: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			fmt.Println("Please enter Rock, Paper, or Scissors.")
			fmt.Println("1: Rock")
			fmt.Println("2: Paper")
			fmt.Println("3: Scissors")
			fmt.Print("please enter >")

			scanner.Scan()
			in = scanner.Text()
			switch in {
			case "1", "2", "3":
				handShapes, _ := strconv.Atoi(in)
				service.PlayGame(int32(handShapes))
			default:
				fmt.Println("Invalid command.")
			}
		case "2":
			fmt.Println("Here are your match results.")
			service.ReportMatchResults()
		case "3":
			fmt.Println("bye.")
			is_finished = true
		default:
			fmt.Println("Invalid command.")
		}
	}
}
