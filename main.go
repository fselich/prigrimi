package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u", "á", "é", "í", "ó", "ú"}

type model struct {
	textInput textinput.Model
	err       error
}

type (
	errMsg error
)

func main() {
	fmt.Print("\033[h\033[2J")
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Hello..."
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 200

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"  \033[33mWrite the text you want to annoy with\u001B[0m\n\n  %s\n  > \u001B[32m%s\u001B[0m",
		m.textInput.View(),
		convert(m.textInput.Value()),
	) + "\n"
}

func convert(text string) string {
	for _, vowel := range vowels {
		lowerVowel := strings.ToLower(vowel)
		upperVowel := strings.ToUpper(vowel)
		text = strings.ReplaceAll(text, lowerVowel, "i")
		text = strings.ReplaceAll(text, upperVowel, "I")
	}

	return text
}
