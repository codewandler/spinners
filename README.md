<div align="center">

# ✦ spinners

**21 width-adaptive animated spinners for terminal UIs**

Built as a drop-in [Bubble Tea](https://charm.land/bubbletea) component.
Every spinner renders into exactly `width` columns — resize at any time, no reconfiguration needed.

[![Go Reference](https://pkg.go.dev/badge/github.com/codewandler/spinners.svg)](https://pkg.go.dev/github.com/codewandler/spinners)
[![Go Report Card](https://goreportcard.com/badge/github.com/codewandler/spinners)](https://goreportcard.com/report/github.com/codewandler/spinners)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

<br>

![Demo](demo.gif)

</div>

---

## ✨ Features

- **21 built-in spinners** — from Knight Rider to Binary Rain, each with a distinct visual style
- **Width-adaptive** — pass any column count; the spinner fills it exactly, every frame
- **Bubble Tea native** — implements the standard `Init / Update / View` pattern as a composable sub-model
- **Runtime flexibility** — swap spinner, resize, or change background color on the fly
- **Zero dependencies** beyond Bubble Tea and lipgloss

---

## 📦 Install

```bash
go get github.com/codewandler/spinners@latest
```

---

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "os"

    tea "charm.land/bubbletea/v2"
    "github.com/codewandler/spinners"
)

type model struct {
    spinner spinners.Model
}

func (m model) Init() tea.Cmd {
    return m.spinner.Tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyPressMsg:
        return m, tea.Quit
    }
    var cmd tea.Cmd
    m.spinner, cmd = m.spinner.Update(msg)
    return m, cmd
}

func (m model) View() tea.View {
    return tea.NewView(m.spinner.View())
}

func main() {
    m := model{spinner: spinners.New(spinners.KnightRider, 20)}
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

---

## 🎨 Spinner Catalog

All spinners are available via `spinners.All` or by name (e.g. `spinners.KnightRider`).

| # | Name | Description |
|---|------|-------------|
| 1 | **Knight Rider** | Blue glow bouncing left↔right with halo falloff |
| 2 | **Plasma Pulse** | Violet/purple wave scrolling across the bar |
| 3 | **DNA Helix** | Paired complementary dots in two interlaced strands |
| 4 | **Matrix** | Cascading katakana/digit rain with brightening columns |
| 5 | **Equalizer** | Bars bouncing at independent phases like an audio visualizer |
| 6 | **Comet** | Glowing head with multi-shade trail sweeping across |
| 7 | **Lava Lamp** | Amber blobs pulsing at staggered phases |
| 8 | **Aurora** | Rippling teal/green shimmer bands |
| 9 | **Glitch** | Corrupted data bar with deterministic per-cell noise |
| 10 | **Warp Drive** | Concentric rings rushing toward the viewer |
| 11 | **Neon Bounce** | Rainbow gem bouncing inside the bar with trailing dots |
| 12 | **Heartbeat** | ECG-style pulse spike traveling across a flat baseline |
| 13 | **Scan Line** | Bright radar beam sweeping through a dark dot grid |
| 14 | **Tidal Wave** | Sine-shaped water crest rolling with ocean colors |
| 15 | **Starfield** | Sparse stars at varying brightness drifting through deep space |
| 16 | **Serpentine** | Smooth block-width ribbon cycling the rainbow as it scrolls |
| 17 | **Quantum Foam** | Cells flicker between particle states in pulsing energy waves |
| 18 | **Thunderstrike** | Cells independently flash white/yellow then fade through purple |
| 19 | **Ice Crystals** | Frost radiates outward from the center using angular chars |
| 20 | **Solar Flare** | Roiling red surface with staggered eruptions decaying orange→white |
| 21 | **Binary Rain** | 0s and 1s at different per-column speeds, brightening as they stream |

---

## 📖 API

### Types

```go
// Spinner is a named, width-adaptive animation.
type Spinner struct {
    Name   string
    Frames func(tick, width int) string
    FPS    time.Duration
}

// Model is an embeddable Bubble Tea component.
type Model struct { ... }
```

### Functions

| Function | Description |
|----------|-------------|
| `New(s Spinner, width int) Model` | Create a new Model for the given spinner and column width |
| `(m *Model) SetWidth(w int)` | Change render width at runtime |
| `(m Model) Width() int` | Get current render width |
| `(m *Model) SetSpinner(s Spinner)` | Swap spinner, reset tick to 0 |
| `(m *Model) SetBackground(color string)` | Set background color (hex, e.g. `"#141414"`) |
| `(m Model) Tick() tea.Cmd` | Start the animation loop — call from `Init` |
| `(m Model) Update(msg tea.Msg) (Model, tea.Cmd)` | Handle tick messages |
| `(m *Model) View() string` | Render current frame into exactly `width` columns |

### Iterating all spinners

```go
for _, sp := range spinners.All {
    fmt.Println(sp.Name, sp.Frames(0, 20))
}
```

---

## 🖥️ Demo

Run the interactive showcase locally:

```bash
go run ./cmd/demo
# or with Task:
task demo
```

To regenerate `demo.gif`:

```bash
task gif
# (requires Docker — see Taskfile.yaml)
```

---

## 📄 License

[MIT](LICENSE) © codewandler
