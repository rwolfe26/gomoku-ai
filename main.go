package main

import "fmt"

func main() {
	board := NewBoard(9)
	var r, c int
	player := 'X'

	for {
		board.Print()

		if player == 'X' {
			fmt.Printf("Player %c, enter row and column: ", player)

			n, err := fmt.Scanln(&r, &c)
			if err != nil || n != 2 {
				fmt.Println("Please enter two integers like: 0 2")
				continue
			}

			err = board.PlaceStone(r, c, player)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			// move := BestMoveOnePly(board, 'O')
			move := BestMoveMinimax(board, 'O', 2)
			fmt.Printf("Computer chooses: %d %d\n", move.row, move.col)
			_ = board.PlaceStone(move.row, move.col, 'O')
		}

		if board.HasFive(player) {
			board.Print()
			fmt.Printf("Player %c wins!\n", player)
			break
		}

		fmt.Println("Evaluation for X:", Evaluate(board, 'X'))
		fmt.Println("Evaluation for O:", Evaluate(board, 'O'))

		if board.IsFull() {
			board.Print()
			fmt.Println("It's a draw!")
			break
		}

		if player == 'X' {
			player = 'O'
		} else {
			player = 'X'
		}
	}
}
