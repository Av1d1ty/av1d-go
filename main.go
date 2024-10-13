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
    viewProps viewProps
	textInput textinput.Model
	engine    engine.Engine
	board     [][]int
	err       error
}

type viewProps struct {
	header     string
	padNumbers bool
}

func initialModel() model {
	boardSize := engine.Board9x9 // TODO: get from user
	ti := textinput.New()
	ti.Placeholder = "a1" // TODO: Remove after first run
	ti.Focus()
	viewProps := viewProps{header: getHeader(boardSize)}

	if boardSize > engine.Board9x9 {
		viewProps.padNumbers = true
		ti.CharLimit = 3
	} else {
		ti.CharLimit = 2
	}

	e := engine.Engine{}
	board := e.InitBoard(boardSize, 1, 2, true)

	return model{
		textInput: ti,
		engine:    e,
		board:     board,
        viewProps: viewProps,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			m.err = m.engine.MakeMove(m.textInput.Value())
			m.textInput.Reset()
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var sb strings.Builder

	sb.WriteString(m.viewProps.header)
	for i, row := range m.board {
		if m.viewProps.padNumbers && i < 9 {
			sb.WriteString(fmt.Sprintf("\n %d %v", i+1, row))
		} else {
			sb.WriteString(fmt.Sprintf("\n%d %v", i+1, row))
		}
	}

	if m.err != nil {
		return fmt.Sprintf(
			"%s\n\n%s\n%s\n%s",
			sb.String(), m.textInput.View(), m.err,
			"(esc to quit)",
		) + "\n"
	}
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		sb.String(), m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func getHeader(boardSize engine.BoardSize) string {
	var sb strings.Builder
    if boardSize <= engine.Board9x9{
        sb.WriteString("  ")
    } else {
        sb.WriteString("   ")
    }
	for i := range boardSize {
		sb.WriteString(fmt.Sprintf(" %c", 97+i))
	}
	return sb.String()
}
