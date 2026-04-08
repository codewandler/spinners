// Package main is the spinner showcase demo.
// Run with: go run ./cmd/demo
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/codewandler/spinners"
)

// demoWidths cycles through these values every ~2 s to show width-adaptability.
var demoWidths = []int{8, 12, 16, 20, 24}

type tickMsg struct{ t time.Time }

type model struct {
	tick     int
	quitting bool
}

func (m model) Init() tea.Cmd { return nextTick() }

func nextTick() tea.Cmd {
	return tea.Tick(time.Second/16, func(t time.Time) tea.Msg { return tickMsg{t} })
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
	case tickMsg:
		m.tick++
		return m, nextTick()
	}
	return m, nil
}

func (m model) currentWidth() int {
	return demoWidths[(m.tick/32)%len(demoWidths)]
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#7C3AED")).
			Padding(0, 2)

	nameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8")).
			Width(16)

	widthPillStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F59E0B")).
			Bold(true)

	boxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#374151")).
			Padding(1, 2)

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4B5563"))
)

func (m model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	w := m.currentWidth()

	var sb strings.Builder
	sb.WriteString(titleStyle.Render("✦  Spinner Showcase") + "\n")
	sb.WriteString("   " + dimStyle.Render("width = ") + widthPillStyle.Render(fmt.Sprintf("%d", w)) +
		dimStyle.Render("  (cycles automatically)") + "\n\n")

	all := spinners.All
	cols := 2
	rows := (len(all) + cols - 1) / cols

	cells := make([]string, len(all))
	for i, sp := range all {
		num := fmt.Sprintf("%2d. ", i+1)
		label := nameStyle.Render(sp.Name)
		frame := sp.Frames(m.tick, w)
		cells[i] = dimStyle.Render(num) + label + "  " + frame
	}

	colWidth := 46
	for row := 0; row < rows; row++ {
		var rowParts []string
		for col := 0; col < cols; col++ {
			idx := row*cols + col
			if idx >= len(cells) {
				break
			}
			rowParts = append(rowParts, lipgloss.PlaceHorizontal(colWidth, lipgloss.Left, cells[idx]))
		}
		sb.WriteString("  " + strings.Join(rowParts, "  ") + "\n")
	}

	sb.WriteString("\n")
	sb.WriteString(boxStyle.Render(
		dimStyle.Render("press ") +
			lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")).Bold(true).Render("q") +
			dimStyle.Render(" to quit"),
	))
	sb.WriteString("\n")

	return tea.NewView(sb.String())
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
