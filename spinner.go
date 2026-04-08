// Package spinners provides a collection of width-adaptive animated spinners
// and a reusable Bubble Tea component for embedding them in TUI applications.
//
// Every [Spinner] renders into exactly `width` terminal columns, so a single
// spinner instance can be used at any size without reconfiguration.
//
// Typical embedding:
//
//	m.spinner = spinners.New(spinners.KnightRider, 16)
//	// in Init:  return m.spinner.Tick()
//	// in Update: m.spinner, cmd = m.spinner.Update(msg)
//	// in View:  m.spinner.View()
package spinners

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// ---------------------------------------------------------------------------
// Core types
// ---------------------------------------------------------------------------

// Spinner is a named, width-adaptive animation.
// Frames(tick, width) must return a string whose visible terminal width equals
// the width argument — callers rely on this for layout alignment.
type Spinner struct {
	// Name is a human-readable label shown in the showcase.
	Name string
	// Frames renders one animation frame at the given tick counter and width.
	Frames func(tick, width int) string
	// FPS is the recommended refresh rate. Defaults to 16 fps if zero.
	FPS time.Duration
}

// fps returns the effective frame rate.
func (s Spinner) fps() time.Duration {
	if s.FPS > 0 {
		return s.FPS
	}
	return time.Second / 16
}

// ---------------------------------------------------------------------------
// Reusable Bubble Tea model
// ---------------------------------------------------------------------------

// TickMsg is the internal tick message used by [Model].
type TickMsg struct{ id int }

// Model is an embeddable Bubble Tea component that drives a single [Spinner].
type Model struct {
	spinner    Spinner
	width      int
	tick       int
	id         int    // disambiguates ticks when multiple Models run simultaneously
	bg         string // optional background color (e.g. "#141414"); empty means no background
	bgStyle    lipgloss.Style // cached whitespace style — rebuilt only when bg changes
	bgStyleSet bool           // true once bgStyle has been built for the current bg value
}

var modelIDCounter int

// New creates a Model for the given spinner at the given column width.
func New(s Spinner, width int) Model {
	modelIDCounter++
	return Model{spinner: s, width: width, id: modelIDCounter}
}

// SetWidth changes the render width at runtime.
func (m *Model) SetWidth(w int) { m.width = w }

// Width returns the current render width.
func (m Model) Width() int { return m.width }

// Spinner returns the underlying Spinner definition.
func (m Model) Spinner() Spinner { return m.spinner }

// SetSpinner swaps the active spinner definition and resets the tick counter
// to zero so the new animation starts from its first frame. The model id is
// preserved, so the existing tick-loop (started by a prior Tick() call)
// continues driving the animation without interruption.
func (m *Model) SetSpinner(s Spinner) { m.spinner = s; m.tick = 0 }

// SetBackground sets the background color applied behind each rendered frame.
// Pass an empty string to clear the background.
func (m *Model) SetBackground(color string) { m.bg = color }

// Background returns the current background color.
func (m Model) Background() string { return m.bg }

// Tick returns the initial command that starts the animation loop.
// Call this from your model's Init method (or Batch it in).
func (m Model) Tick() tea.Cmd {
	return tea.Tick(m.spinner.fps(), func(_ time.Time) tea.Msg {
		return TickMsg{id: m.id}
	})
}

// Update handles TickMsg and advances the animation.
// Returns the updated Model and the next tick command.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if t, ok := msg.(TickMsg); ok && t.id == m.id {
		m.tick++
		return m, m.Tick()
	}
	return m, nil
}

// View renders the current frame into exactly m.Width() terminal columns.
// A minimum of 2 is enforced so bounce-style spinners (which compute
// (w-1)*2 as a period) never divide by zero.
// If a background color has been set via SetBackground, it is applied as the
// cell background for the entire rendered line.
func (m *Model) View() string {
	w := m.width
	if w < 2 {
		w = 2
	}
	frame := m.spinner.Frames(m.tick, w)
	if m.bg != "" {
		if !m.bgStyleSet {
			m.bgStyle = lipgloss.NewStyle().Background(lipgloss.Color(m.bg))
			m.bgStyleSet = true
		}
		frame = lipgloss.PlaceHorizontal(
			w,
			lipgloss.Left,
			frame,
			lipgloss.WithWhitespaceStyle(m.bgStyle),
		)
	}
	return frame
}

// ---------------------------------------------------------------------------
// Helpers (unexported, used by spinner frame functions)
// ---------------------------------------------------------------------------

func iabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func iclamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
