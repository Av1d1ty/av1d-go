package engine

import (
	"fmt"
	"strconv"
)

type Engine struct {
	board      [][]int
	boardSize  int
	playerInt  int
    machineInt int
	playerTurn bool
}

type move struct {
	row int
	col int
}

func (e *Engine) InitBoard(size int, playerInt int, machineInt int, playerTurn bool) [][]int {
    e.boardSize = size
    e.playerInt = playerInt
    e.machineInt = machineInt
    e.playerTurn = playerTurn

    e.board = make([][]int, size)
    for i := range e.board {
        e.board[i] = make([]int, size)
    }
    return e.board
}

// Parse input string and make the move
func (e *Engine) MakeMove(input string) error {
	move, err := e.getMoveFromString(input)
	if err != nil {
		return err
	}
	e.makeMove(move)
	return nil
}

func (e *Engine) getMoveFromString(input string) (move, error) {
    inputLen := len(input)
	if inputLen < 2 {
		return move{}, fmt.Errorf("Input is too short")
	}
	if inputLen > 3 || (e.boardSize < 10 && inputLen > 2) {
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
    return move{row: row-1, col: int(col-97)}, nil
}

func (m *Engine) makeMove(move move) {
	m.board[move.row][move.col] = m.playerInt
	// TODO: further game logic
}
