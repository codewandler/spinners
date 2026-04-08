Output demo.gif

# Terminal dimensions — 96 cols fits both spinner columns exactly, 22 rows gives
# breathing room around the 18-line output.
Set Width  1000
Set Height 500
Set FontSize 15
Set FontFamily "JetBrains Mono"
Set Framerate 15
Set Theme "Dracula"

# Binary is pre-built on the host (CGO_ENABLED=0 go build -o spinners-demo ./cmd/demo)
# and mounted into the container at /vhs/spinners-demo.
Hide
Type "clear"
Enter
Show

Type "./spinners-demo"
Enter

# One full width-cycle: 5 widths × 2 s each = 10 s.
# Start just after tick 0 so the first width (8) is already visible, then watch
# 12 → 16 → 20 → 24 → back to 8.
Sleep 11s

Type "q"
Sleep 0.5

Hide
Type "rm spinners-demo"
Enter
