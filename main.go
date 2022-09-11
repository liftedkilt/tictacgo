package main

import (
	"fmt"
	"os"

	"github.com/liftedkilt/tictactoe/tictactoe"
)

func main() {
	b := tictactoe.New()

	play(b)
	// p := tictactoe.Position{
	// 	Col: 2,
	// 	Row: 2,
	// }

	// b = b.Mark(p, tictactoe.O)

}

func play(b tictactoe.Board) {
	var input string

	fmt.Println(b)

	for {
		// b.BestMove(tictactoe.X)

		fmt.Printf("Enter your coordinates: ")

		fmt.Scanln(&input)

		p, err := tictactoe.ParsePosition(input)

		if err != nil {
			fmt.Println("Please enter coordinates in form: A2, C0, etc...")

			continue
		}

		// b, err = b.Mark(p, tictactoe.X)

		// if err != nil {
		// 	fmt.Println("Chosen space already occupied!")

		// 	continue
		// }
		b[p.Row][p.Col] = tictactoe.X

		if b.IsEndState() {
			fmt.Println("You have won!")

			fmt.Println(b)

			os.Exit(0)
		}

		pos := b.BestMove(tictactoe.O)

		fmt.Println(pos.Row+1, pos.Col+1)

		b[pos.Row][pos.Col] = tictactoe.O

		fmt.Println("AI has played:")

		fmt.Println(b)

		if b.IsEndState() {
			fmt.Println("AI has won!")

			fmt.Println(b)

			os.Exit(0)
		}
	}
}
