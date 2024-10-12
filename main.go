package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Av1d1ty/av1d-go/engine"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type errMsg error

type model struct {
	textInput textinput.Model
	engine    engine.Engine
	board     [][]int
	err       error
}

func initialModel() model {
    boardSize := 9 // TODO: get from user
	ti := textinput.New()
    ti.Placeholder = "a1" // TODO: Remove after first run
	ti.Focus()
	ti.CharLimit = len(fmt.Sprintf("%v", boardSize)) + 1
	ti.Width = 5
	e := engine.Engine{}
	board := e.InitBoard(boardSize, 1, 2, true)

	return model{
		textInput: ti,
		engine:    e,
		board:     board,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
        case "enter":
            m.err = m.engine.MakeMove(m.textInput.Value())    
            m.textInput.Reset()
		}
	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// TODO: make view proper when size > 9
	var rowStrings []string
	rowStrings = append(rowStrings, "   a b c d e f g h i")
	for i, row := range m.board {
		rowStrings = append(rowStrings, fmt.Sprintf("%d %v", i+1, row))
	}
    
    if m.err != nil {
        // return fmt.Sprintf("\nError: %v\n", m.err)
        return fmt.Sprintf(
            "%v\n\n%s\n%v\n%s",
            strings.Join(rowStrings, "\n"), m.textInput.View(),
            m.err,
            "(esc to quit)",
        ) + "\n"
    }
	return fmt.Sprintf(
		"%v\n\n%s\n\n%s",
		strings.Join(rowStrings, "\n"), m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
