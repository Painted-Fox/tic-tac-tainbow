package main

import (
	"fmt"
)

/*
We map the board into fields numbered from 1 to 9, then we map these fields to
an unsigned thirty-two bit integer like so...

  1  |  2  |  3
 ---------------
  4  |  5  |  6
 ---------------
  7  |  8  |  9

Field      9| 8| 7| 6| 5| 4| 3| 2| 1| Winner
uint32 ...00|00|00|00|00|00|00|00|00|00

The first two bits are reserved for when a player has won the game.  The first
bit is set when X has won.  The second bit is set when O has won.  If the board
is in an invalid state, both the first and second bit will be set.

After the first two bits, each field on the game board is represented by two
bits.  The first bit is set if that field is occupied by X and the second bit
is set if that field is occupied by O.

In this version of Tic-Tac-Toe, we assume that X is always the first player.
*/

const (
	_ = iota
	X
	O
	ERR
)

type TicTacToeBoard uint32

func (board TicTacToeBoard) SetX(pos uint32) TicTacToeBoard {
	return TicTacToeBoard(uint32(board) | 1<<(2*pos)).setWinner()
}

func (board TicTacToeBoard) SetO(pos uint32) TicTacToeBoard {
	return TicTacToeBoard(uint32(board) | 1<<(2*pos+1)).setWinner()
}

func (board TicTacToeBoard) GetPos(pos uint32) string {
	field := uint32(board) >> (2 * pos) & 3
	switch field {
	default:
		return " "
	case X:
		return "X"
	case O:
		return "O"
	case ERR:
		return "!"
	}

}

func (board TicTacToeBoard) GetWinner() string {
	win := uint32(board) & 3
	switch win {
	default:
		return " "
	case X:
		return "X"
	case O:
		return "O"
	case ERR:
		return "!"
	}
}

// TODO
func (board TicTacToeBoard) setWinner() TicTacToeBoard {
	var i uint32
	for i = 0; i < 3; i++ {
		if (uint32(board)>>(i*2+2))&21 == 21 {
			return TicTacToeBoard(uint32(board) | 1)
		}
		if (uint32(board)>>(i*2+2))&42 == 42 {
			return TicTacToeBoard(uint32(board) | 2)
		}
	}

	return board
}

func (board *TicTacToeBoard) String() string {
	return fmt.Sprintf(` %s  |  %s  |  %s
 -------------
 %s  |  %s  |  %s
 -------------
 %s  |  %s  |  %s`,
		board.GetPos(1),
		board.GetPos(2),
		board.GetPos(3),
		board.GetPos(4),
		board.GetPos(5),
		board.GetPos(6),
		board.GetPos(7),
		board.GetPos(8),
		board.GetPos(9))
}

func main() {
	var board TicTacToeBoard
	board = board.SetX(1)
	board = board.SetO(5)
	board = board.SetX(9)
	board = board.SetO(6)
	board = board.SetX(2)
	board = board.SetO(4)

	fmt.Printf("%s\n\n", board.String())

	fmt.Printf("%d\n", uint32(board))
	fmt.Printf("%s\n", board.GetWinner())
}
