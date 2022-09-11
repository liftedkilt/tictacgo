package tictactoe

import (
	"math"
)

func (b Board) BestMove(player Marker) Position {
	// AI to make its turn
	bestScore := math.Inf(-1)
	move := Position{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Is the spot available?
			if b[i][j] == B {
				b[i][j] = player
				score := minimax(b, player, 0)
				b[i][j] = B
				if score > bestScore {
					bestScore = score
					move = Position{Row: i, Col: j}
				}
			}
		}
	}
	return move
}

func minimax(b Board, player Marker, depth int) float64 {
	if b.IsEndState() {
		switch player {
		case X:
			return -10 + float64(depth)
		case O:
			return 10 - float64(depth)
		default:
			panic("unreachable")
		}
	}

	var (
		compfunc   func(x float64, y float64) float64
		sign       int
		nextPlayer Marker
	)

	if player == O {
		compfunc = math.Max
		sign = -1
		nextPlayer = X
	} else {
		compfunc = math.Min
		sign = 1
		nextPlayer = O
	}

	bestVal := math.Inf(sign)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == " " {
				b[i][j] = player

				value := minimax(b, nextPlayer, depth+1)

				b[i][j] = " "

				bestVal = compfunc(bestVal, value)
			}
		}
	}
	return bestVal
}
