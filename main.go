package main

import "fmt"

func main() {
	board := NewBoard(9)
	var r, c int
	player := 'X'

	for {
		fmt.Printf("Player %c, enter row and column: ", player)

		n, err := fmt.Scanln(&r, &c)
		if err != nil || n != 2 {
			fmt.Println("Please enter two integers likeL: 0 2")
			continue
		}

		err = board.PlaceStone(r, c, player)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if board.HasFive(player) {
			board.Print()
			fmt.Printf("Player %c wins!\n", player)
			break
		}

		board.Print()

		if player == 'X' {
			player = 'O'
		} else {
			player = 'X'
		}

	}
}
