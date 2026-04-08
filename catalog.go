package spinners

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// ---------------------------------------------------------------------------
// Spinner catalog
// ---------------------------------------------------------------------------

// All is the ordered list of every built-in spinner.
var All = []Spinner{
	KnightRider,
	PlasmaPulse,
	DNAHelix,
	Matrix,
	Equalizer,
	Comet,
	LavaLamp,
	Aurora,
	Glitch,
	WarpDrive,
	NeonBounce,
	Heartbeat,
	ScanLine,
	TidalWave,
	Starfield,
	Serpentine,
	QuantumFoam,
	Thunderstrike,
	IceCrystals,
	SolarFlare,
	BinaryRain,
}

// Pre-built per-color styles for every catalog spinner.
// These are package-level vars so they are constructed once at program start
// rather than being re-allocated on every Frames() tick.
var (
	// KnightRider
	knightRiderHead = lipgloss.NewStyle().Foreground(lipgloss.Color("#60A5FA")).Bold(true)
	knightRiderH1   = lipgloss.NewStyle().Foreground(lipgloss.Color("#1D4ED8"))
	knightRiderH2   = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E3A5F"))
	knightRiderDark = lipgloss.NewStyle().Foreground(lipgloss.Color("#0F172A"))

	// PlasmaPulse
	plasmaPulseStyles = buildSpinnerStyles([]string{
		"#4C1D95", "#6D28D9", "#7C3AED", "#8B5CF6",
		"#A78BFA", "#8B5CF6", "#7C3AED", "#6D28D9",
	})

	// DNAHelix
	dnaTop = lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981"))
	dnaBot = lipgloss.NewStyle().Foreground(lipgloss.Color("#3B82F6"))

	// Matrix
	matrixStyles = buildSpinnerStyles([]string{
		"#14532D", "#166534", "#15803D", "#16A34A", "#22C55E", "#4ADE80",
	})

	// Equalizer
	equalizerStyles = buildSpinnerStyles([]string{
		"#6D28D9", "#7C3AED", "#8B5CF6", "#A78BFA", "#C4B5FD",
	})

	// Comet
	cometHeadStyles  = buildBoldSpinnerStyles([]string{"#FFFFFF", "#E0F2FE", "#BAE6FD", "#7DD3FC", "#38BDF8"})
	cometTrailStyles = buildSpinnerStyles([]string{"#0EA5E9", "#0284C7", "#0369A1", "#075985", "#0C4A6E"})

	// LavaLamp
	lavaLampStyles = buildSpinnerStyles([]string{
		"#92400E", "#B45309", "#D97706", "#F59E0B", "#FBBF24", "#FDE68A",
	})

	// Aurora
	auroraStyles = buildSpinnerStyles([]string{
		"#042f2e", "#065f46", "#047857", "#059669",
		"#34d399", "#6ee7b7", "#a7f3d0", "#6ee7b7",
		"#34d399", "#059669", "#047857", "#065f46",
	})

	// Glitch
	glitchStyles = buildSpinnerStyles([]string{
		"#EF4444", "#F87171", "#FECACA",
		"#E5E7EB", "#9CA3AF", "#6B7280",
		"#EF4444", "#FFFFFF", "#374151", "#EF4444",
	})

	// WarpDrive
	warpDriveStyles = buildSpinnerStyles([]string{
		"#1E3A5F", "#1D4ED8", "#2563EB", "#3B82F6",
		"#60A5FA", "#93C5FD", "#DBEAFE",
		"#93C5FD", "#60A5FA", "#3B82F6", "#2563EB",
	})

	// NeonBounce
	neonBounceTrack = lipgloss.NewStyle().Foreground(lipgloss.Color("#1F2937"))
	neonBounceDot   = lipgloss.NewStyle().Foreground(lipgloss.Color("#6B21A8"))
	neonBounceHues  = buildBoldSpinnerStyles([]string{
		"#F43F5E", "#FB923C", "#FACC15", "#4ADE80", "#38BDF8", "#818CF8", "#E879F9",
	})

	// Heartbeat
	heartbeatDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4B0017"))
	heartbeatStyles = buildSpinnerStyles([]string{
		"#BE123C", "#E11D48", "#F43F5E", "#FB7185",
		"#FECDD3", "#FB7185", "#F43F5E", "#E11D48",
	})

	// ScanLine
	scanLineBg     = lipgloss.NewStyle().Foreground(lipgloss.Color("#0F172A"))
	scanLineStyles = buildSpinnerStyles([]string{"#FFFFFF", "#CBD5E1", "#64748B", "#1E293B"})

	// TidalWave
	tidalWaveStyles = buildSpinnerStyles([]string{
		"#0C4A6E", "#075985", "#0369A1", "#0284C7",
		"#0EA5E9", "#38BDF8", "#7DD3FC", "#BAE6FD",
	})

	// Starfield
	starfieldDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("#0F172A"))
	starfieldStyles = buildSpinnerStyles([]string{
		"#1E293B", "#334155", "#475569", "#94A3B8", "#E2E8F0", "#FFFFFF",
	})

	// Serpentine
	serpentineStyles = buildSpinnerStyles([]string{
		"#F43F5E", "#F97316", "#FACC15", "#4ADE80",
		"#38BDF8", "#818CF8", "#E879F9", "#F43F5E",
	})

	// QuantumFoam
	quantumCoolStyles = buildSpinnerStyles([]string{"#1E3A5F", "#1D4ED8", "#2563EB", "#3B82F6"})
	quantumHotStyles  = buildSpinnerStyles([]string{"#7DD3FC", "#BAE6FD", "#E0F2FE", "#FFFFFF"})

	// Thunderstrike
	thunderFlashBold   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFBEB")).Bold(true)
	thunderFlash       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FBBF24"))
	thunderAfterStyles = buildSpinnerStyles([]string{"#7C3AED", "#6D28D9", "#4C1D95", "#2E1065"})
	thunderDark        = lipgloss.NewStyle().Foreground(lipgloss.Color("#0F172A"))

	// IceCrystals
	iceDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("#0F172A"))
	iceStyles = buildSpinnerStyles([]string{
		"#0C4A6E", "#075985", "#0369A1", "#0284C7",
		"#38BDF8", "#7DD3FC", "#BAE6FD", "#E0F2FE",
	})

	// SolarFlare
	solarFlareBaseStyles  = buildSpinnerStyles([]string{"#7F1D1D", "#991B1B", "#B91C1C", "#DC2626"})
	solarFlareStyles      = buildBoldSpinnerStyles([]string{"#EF4444", "#F97316", "#FBBF24", "#FEF9C3", "#FFFFFF"})
	solarFlareDecayStyles = buildSpinnerStyles([]string{"#EF4444", "#F97316", "#FBBF24", "#FEF9C3", "#FFFFFF"})

	// BinaryRain
	binaryRainStyles = buildSpinnerStyles([]string{
		"#14532D", "#15803D", "#16A34A", "#22C55E", "#4ADE80", "#86EFAC",
	})
)

// buildSpinnerStyles creates a slice of pre-built foreground styles for the
// given hex colors. Slices are constructed once at package init so spinner
// Frames() closures never allocate lipgloss.Style on the hot render path.
func buildSpinnerStyles(colors []string) []lipgloss.Style {
	s := make([]lipgloss.Style, len(colors))
	for i, c := range colors {
		s[i] = lipgloss.NewStyle().Foreground(lipgloss.Color(c))
	}
	return s
}

// buildBoldSpinnerStyles is like buildSpinnerStyles but marks each style Bold.
func buildBoldSpinnerStyles(colors []string) []lipgloss.Style {
	s := make([]lipgloss.Style, len(colors))
	for i, c := range colors {
		s[i] = lipgloss.NewStyle().Foreground(lipgloss.Color(c)).Bold(true)
	}
	return s
}

// KnightRider — blue glow bouncing left↔right with halo falloff.
var KnightRider = Spinner{
	Name: "Knight Rider",
	Frames: func(tick, w int) string {
		totalFrames := (w - 1) * 2
		pos := tick % totalFrames
		if pos >= w {
			pos = totalFrames - pos
		}
		bar := make([]string, w)
		for i := range bar {
			dist := iabs(i - pos)
			switch dist {
			case 0:
				bar[i] = knightRiderHead.Render("█")
			case 1:
				bar[i] = knightRiderH1.Render("▓")
			case 2:
				bar[i] = knightRiderH2.Render("▒")
			default:
				bar[i] = knightRiderDark.Render("░")
			}
		}
		return strings.Join(bar, "")
	},
}

// PlasmaPulse — violet/purple wave scrolling across the bar.
var PlasmaPulse = Spinner{
	Name: "Plasma Pulse",
	Frames: func(tick, w int) string {
		n := len(plasmaPulseStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			sb.WriteString(plasmaPulseStyles[(tick+i)%n].Render("█"))
		}
		return sb.String()
	},
}

// DNAHelix — paired complementary dots scrolling in two interlaced strands.
var DNAHelix = Spinner{
	Name: "DNA Helix",
	Frames: func(tick, w int) string {
		top := []string{"·", "·", "◦", "○", "◦", "·"}
		bot := []string{"○", "◦", "·", "·", "·", "◦"}
		n := len(top)
		pairs := w / 2
		var sb strings.Builder
		for i := 0; i < pairs; i++ {
			idx := (tick + i) % n
			sb.WriteString(dnaTop.Render(top[idx]) + dnaBot.Render(bot[(idx+n/2)%n]))
		}
		if w%2 != 0 {
			sb.WriteString(dnaTop.Render("·"))
		}
		return sb.String()
	},
}

// Matrix — cascading katakana/digit rain with brightening columns.
var Matrix = Spinner{
	Name: "Matrix",
	Frames: func(tick, w int) string {
		chars := []string{"0", "1", "ﾊ", "ﾐ", "ﾋ", "ｰ", "ｳ", "ｼ", "ﾅ", "ﾓ", "ﾆ", "ｻ"}
		nc := len(chars)
		nb := len(matrixStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			c := chars[(tick*3+i*7)%nc]
			sb.WriteString(matrixStyles[(tick+i)%nb].Render(c))
		}
		return sb.String()
	},
}

// Equalizer — N bars bouncing at independent phases like an audio visualizer.
var Equalizer = Spinner{
	Name: "Equalizer",
	Frames: func(tick, w int) string {
		heights := []string{"▁", "▃", "▅", "▇"}
		nh := len(heights)
		nc := len(equalizerStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			phase := (i * 3) % 7
			h := (tick + phase) % (nh * 2)
			if h >= nh {
				h = nh*2 - 1 - h
			}
			sb.WriteString(equalizerStyles[i%nc].Render(heights[h]))
		}
		return sb.String()
	},
}

// Comet — glowing head with multi-shade trail sweeping across the bar.
var Comet = Spinner{
	Name: "Comet",
	Frames: func(tick, w int) string {
		tailLen := iclamp(w/3, 2, 6)
		tailChars := []string{"█", "▓", "▒", "░", "·", " "}
		pos := tick % (w + tailLen)
		cells := make([]string, w)
		for i := range cells {
			cells[i] = " "
		}
		for j := 0; j < tailLen+1; j++ {
			col := pos - j
			if col < 0 || col >= w {
				continue
			}
			ch := tailChars[iclamp(j, 0, len(tailChars)-1)]
			if j == 0 {
				cells[col] = cometHeadStyles[tick%len(cometHeadStyles)].Render(ch)
			} else {
				cells[col] = cometTrailStyles[iclamp(j-1, 0, len(cometTrailStyles)-1)].Render(ch)
			}
		}
		return strings.Join(cells, "")
	},
}

// LavaLamp — amber blobs pulsing at staggered phases across the bar.
var LavaLamp = Spinner{
	Name: "Lava Lamp",
	Frames: func(tick, w int) string {
		period := 12
		nh := len(lavaLampStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			off := (i * period) / w
			phase := (tick + off) % period
			dist := phase
			if dist > period/2 {
				dist = period - dist
			}
			idx := iclamp(dist, 0, nh-1)
			ch := "●"
			if idx < 2 {
				ch = "·"
			} else if idx < 4 {
				ch = "○"
			}
			sb.WriteString(lavaLampStyles[idx].Render(ch))
		}
		return sb.String()
	},
}

// Aurora — rippling teal/green shimmer bands.
var Aurora = Spinner{
	Name: "Aurora",
	Frames: func(tick, w int) string {
		chars := []string{"░", "▒", "▓", "█", "▓", "▒", "░"}
		nc := len(chars)
		np := len(auroraStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			sb.WriteString(auroraStyles[(tick+i*2)%np].Render(chars[(tick+i)%nc]))
		}
		return sb.String()
	},
}

// Glitch — corrupted data bar with deterministic per-cell noise.
var Glitch = Spinner{
	Name: "Glitch",
	Frames: func(tick, w int) string {
		base := []string{"▓", "░", "▒", "█", "▓", "▒", "░", "▓", "█", "▒"}
		nb := len(base)
		glyphs := []string{"░", "▒", "▓", "█", "▄", "▀", "■", "▪", "▫", "▬"}
		ng := len(glyphs)
		nc := len(glitchStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			seed := (tick*7 + i*13) % 31
			ch := base[i%nb]
			if seed < 8 {
				ch = glyphs[(tick+i*3)%ng]
			}
			sb.WriteString(glitchStyles[(tick+i*4)%nc].Render(ch))
		}
		return sb.String()
	},
}

// WarpDrive — concentric rings rushing toward the viewer.
var WarpDrive = Spinner{
	Name: "Warp Drive",
	Frames: func(tick, w int) string {
		rings := []string{"·", "∘", "○", "◎", "●", "◎", "○", "∘", "·", " ", " "}
		nr := len(rings)
		nc := len(warpDriveStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			idx := (i + tick) % nr
			sb.WriteString(warpDriveStyles[idx%nc].Render(rings[idx]))
		}
		return sb.String()
	},
}

// NeonBounce — rainbow gem bouncing inside the bar with trailing dots.
var NeonBounce = Spinner{
	Name: "Neon Bounce",
	Frames: func(tick, w int) string {
		totalFrames := (w - 1) * 2
		pos := tick % totalFrames
		if pos >= w {
			pos = totalFrames - pos
		}
		cells := make([]string, w)
		for i := range cells {
			cells[i] = neonBounceTrack.Render("─")
		}
		if pos > 0 {
			cells[pos-1] = neonBounceDot.Render("·")
		}
		cells[pos] = neonBounceHues[tick%len(neonBounceHues)].Render("◆")
		if pos < w-1 {
			cells[pos+1] = neonBounceDot.Render("·")
		}
		return strings.Join(cells, "")
	},
}

// Heartbeat — ECG-style pulse spike traveling across a flat baseline.
var Heartbeat = Spinner{
	Name: "Heartbeat",
	Frames: func(tick, w int) string {
		period := w + 4
		beat := tick % period
		levels := []string{"▁", "▂", "▃", "▅", "▇", "█", "▇", "▅", "▃", "▂"}
		nl := len(levels)
		nc := len(heartbeatStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			dist := iabs(i - beat)
			if dist >= nl {
				sb.WriteString(heartbeatDark.Render("▁"))
			} else {
				sb.WriteString(heartbeatStyles[dist%nc].Render(levels[dist]))
			}
		}
		return sb.String()
	},
}

// ScanLine — bright radar beam sweeping right through a dark dot grid.
var ScanLine = Spinner{
	Name: "Scan Line",
	Frames: func(tick, w int) string {
		pos := tick % (w + 4)
		beam := []string{"█", "▓", "▒", "░"}
		bg := scanLineBg.Render("·")
		cells := make([]string, w)
		for i := range cells {
			cells[i] = bg
		}
		for j, bch := range beam {
			col := pos - j
			if col >= 0 && col < w {
				cells[col] = scanLineStyles[j].Render(bch)
			}
		}
		return strings.Join(cells, "")
	},
}

// TidalWave — sine-shaped water crest rolling with ocean colors.
var TidalWave = Spinner{
	Name: "Tidal Wave",
	Frames: func(tick, w int) string {
		wave := []string{"▁", "▂", "▄", "▆", "█", "▆", "▄", "▂", "▁"}
		nw := len(wave)
		nc := len(tidalWaveStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			wavePos := ((i-tick)%nw + nw) % nw
			colorPos := ((i-tick)%nc + nc) % nc
			sb.WriteString(tidalWaveStyles[colorPos].Render(wave[wavePos]))
		}
		return sb.String()
	},
}

// Starfield — sparse stars at varying brightness drifting through deep space.
var Starfield = Spinner{
	Name: "Starfield",
	Frames: func(tick, w int) string {
		stars := []struct{ col, period int }{
			{1, 3}, {3, 5}, {5, 7}, {7, 4},
			{9, 6}, {11, 3}, {13, 5}, {2, 8},
		}
		nd := len(starfieldStyles)
		cells := make([]string, w)
		for i := range cells {
			cells[i] = starfieldDark.Render(" ")
		}
		for _, s := range stars {
			col := ((s.col-tick)%w + w) % w
			if col >= w {
				continue
			}
			brightness := (tick/s.period + s.col) % nd
			ch := "·"
			if brightness >= 3 {
				ch = "✦"
			} else if brightness >= 2 {
				ch = "∗"
			}
			cells[col] = starfieldStyles[brightness].Render(ch)
		}
		return strings.Join(cells, "")
	},
}

// Serpentine — smooth block-width ribbon cycling the rainbow as it scrolls.
var Serpentine = Spinner{
	Name: "Serpentine",
	Frames: func(tick, w int) string {
		band := []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█", "▉", "▊", "▋", "▌", "▍", "▎", "▏"}
		nb := len(band)
		nr := len(serpentineStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			sb.WriteString(serpentineStyles[(i+tick/2)%nr].Render(band[(i+tick)%nb]))
		}
		return sb.String()
	},
}

// QuantumFoam — cells flicker between particle states in pulsing energy waves.
var QuantumFoam = Spinner{
	Name: "Quantum Foam",
	Frames: func(tick, w int) string {
		particles := []string{"·", "∘", "○", "◉", "●", "◉", "○", "∘"}
		np := len(particles)
		nc := len(quantumCoolStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			energy := (tick + i*3) % (np * 2)
			if energy >= np {
				energy = np*2 - 1 - energy
			}
			ch := particles[energy]
			if energy >= np/2 {
				sb.WriteString(quantumHotStyles[(i+tick/3)%nc].Render(ch))
			} else {
				sb.WriteString(quantumCoolStyles[(i+tick/5)%nc].Render(ch))
			}
		}
		return sb.String()
	},
}

// Thunderstrike — cells independently flash white/yellow then fade through purple.
var Thunderstrike = Spinner{
	Name: "Thunderstrike",
	Frames: func(tick, w int) string {
		strikePeriod := 14
		na := len(thunderAfterStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			offset := (i * 5) % strikePeriod
			phase := (tick + offset) % strikePeriod
			switch {
			case phase <= 1:
				sb.WriteString(thunderFlashBold.Render("█"))
			case phase <= 2:
				sb.WriteString(thunderFlash.Render("▇"))
			case phase <= 6:
				sb.WriteString(thunderAfterStyles[iclamp(phase-3, 0, na-1)].Render("▁"))
			default:
				sb.WriteString(thunderDark.Render("░"))
			}
		}
		return sb.String()
	},
}

// IceCrystals — frost radiates outward from the center using angular chars.
var IceCrystals = Spinner{
	Name: "Ice Crystals",
	Frames: func(tick, w int) string {
		glyphs := []string{"·", "╌", "┄", "┈", "╍", "━", "◇", "◈", "❄"}
		ng := len(glyphs)
		ni := len(iceStyles)
		center := w / 2
		period := w + 4
		radius := tick % period
		var sb strings.Builder
		for i := 0; i < w; i++ {
			dist := iabs(i - center)
			if dist <= radius {
				age := radius - dist
				colorIdx := iclamp(age, 0, ni-1)
				glyphIdx := iclamp(age*ng/ni, 0, ng-1)
				sb.WriteString(iceStyles[colorIdx].Render(glyphs[glyphIdx]))
			} else {
				sb.WriteString(iceDark.Render("·"))
			}
		}
		return sb.String()
	},
}

// SolarFlare — roiling red surface with staggered eruptions decaying orange→white.
var SolarFlare = Spinner{
	Name: "Solar Flare",
	Frames: func(tick, w int) string {
		eruptPeriod := 11
		nb := len(solarFlareBaseStyles)
		nf := len(solarFlareStyles)
		chars := []string{"▁", "▂", "▄", "▅", "▆", "▇", "█"}
		nc := len(chars)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			offset := (i * 7) % eruptPeriod
			phase := (tick + offset) % eruptPeriod
			switch {
			case phase <= 1:
				idx := iclamp(nf-1-phase, 0, nf-1)
				sb.WriteString(solarFlareStyles[idx].Render(chars[nc-1-phase]))
			case phase <= 5:
				decay := phase - 2
				sb.WriteString(solarFlareDecayStyles[iclamp(nf-1-decay, 0, nf-1)].Render(chars[iclamp(nc-1-decay, 0, nc-1)]))
			default:
				sb.WriteString(solarFlareBaseStyles[(i+tick/3)%nb].Render(chars[1]))
			}
		}
		return sb.String()
	},
}

// BinaryRain — 0s and 1s at different per-column speeds, brightening as they stream.
var BinaryRain = Spinner{
	Name: "Binary Rain",
	Frames: func(tick, w int) string {
		speeds := []int{3, 5, 2, 7, 4, 6, 3, 5, 2, 4, 7, 3, 6, 2, 5}
		ns := len(speeds)
		nb := len(binaryRainStyles)
		var sb strings.Builder
		for i := 0; i < w; i++ {
			speed := speeds[i%ns]
			colTick := tick * speed
			digit := "0"
			if (colTick+i)%2 == 0 {
				digit = "1"
			}
			bIdx := (colTick/speed + i) % nb
			sb.WriteString(binaryRainStyles[bIdx].Render(digit))
		}
		return sb.String()
	},
}
