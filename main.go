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
	board     [][]engine.Stone
	err       error

	lastBoard    string
	stateChanged bool
}

type viewProps struct {
	header     string
	padNumbers bool
}

var boardTranslation = map[engine.Stone]rune{
	engine.Empty: ' ', engine.Black: '○', engine.White: '●',
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
	board := e.InitBoard(boardSize)

	return model{
		textInput:    ti,
		engine:       e,
		board:        board,
		viewProps:    viewProps,
		stateChanged: true,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.engine.GameEnded {
		return m, tea.Quit
	}

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			m.err = m.engine.MakeMove(m.textInput.Value())
			m.textInput.Reset()
			m.stateChanged = true
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if !m.stateChanged {
		return m.lastBoard
	}
	m.stateChanged = false

	var sb strings.Builder

	sb.WriteString(m.viewProps.header)
	for i, row := range m.board {
		if m.viewProps.padNumbers && i < 9 {
			sb.WriteString("\n ")
		} else {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%d ", i+1))
		for _, item := range row {
			sb.WriteByte(' ')
			sb.WriteRune(boardTranslation[item])
		}
	}
	sb.WriteString(fmt.Sprintf("\n\n%s\n", m.textInput.View()))
	if m.err != nil {
		sb.WriteString(m.err.Error())
	}
	if m.engine.GameEnded {
		sb.WriteString("\nGame over!\n")
	} else {
		sb.WriteString("\nEsc - quit; '/' - pass\n")
	}
	m.lastBoard = sb.String()
	return m.lastBoard
}

func getHeader(boardSize engine.BoardSize) string {
	var sb strings.Builder
	if boardSize <= engine.Board9x9 {
		sb.WriteString("  ")
	} else {
		sb.WriteString("   ")
	}
	for i := range boardSize {
		sb.WriteString(fmt.Sprintf(" %c", 97+i))
	}
	return sb.String()
}
