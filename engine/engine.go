package engine

import (
	"fmt"
	"strconv"
)

type Engine struct {
	board        [][]Stone
	boardSize    BoardSize
	activePlayer Stone
}

type BoardSize int

const (
	Board9x9   BoardSize = 9
	Board13x13 BoardSize = 13
	Board19x19 BoardSize = 19
)

type Stone int

const (
	Empty Stone = 0
	Black Stone = 1
	White Stone = 2
)

type move struct {
	row int
	col int
}

func (e *Engine) InitBoard(size BoardSize) [][]Stone {
	e.boardSize = size
	e.activePlayer = Black

	e.board = make([][]Stone, size)
	for i := range e.board {
		e.board[i] = make([]Stone, size)
	}
	return e.board
}

// Parse input string and make the move
func (e *Engine) MakeMove(input string) error {
	move, err := e.getMoveFromString(input)
	if err != nil {
		return err
	}
	return e.makeMove(move)
}

func (e *Engine) getMoveFromString(input string) (move, error) {
	inputLen := len(input)
	if inputLen < 2 {
		return move{}, fmt.Errorf("Input is too short")
	}
	if inputLen > 3 || (e.boardSize == Board9x9 && inputLen > 2) {
		return move{}, fmt.Errorf("Input is too long")
	}

	col := input[0]
	row, err := strconv.Atoi(input[1:])
	if err != nil {
		return move{}, fmt.Errorf("Invalid row: %s", input[1:])
	}
	// 97 = ASCII 'a'
	if col < 97 || col > byte(97+e.boardSize) {
		return move{}, fmt.Errorf("Invalid column: %c", col)
	}
	return move{row: row - 1, col: int(col - 97)}, nil
}

func (e *Engine) makeMove(move move) error {
	if err := e.validateMove(move); err != nil {
		return err
	}
	e.board[move.row][move.col] = e.activePlayer
    e.revalidateBoard(move)

	if e.activePlayer == Black {
		e.activePlayer = White
	} else {
		e.activePlayer = Black
	}

	return nil
}

func (e *Engine) validateMove(move move) error {
	if e.board[move.row][move.col] != Empty {
		return fmt.Errorf("Position is already taken")
	}
	// TODO: further validatein
	return nil
}

func (e *Engine) revalidateBoard(move move) {
    // TODO: game logic
}
