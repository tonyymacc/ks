# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`ks` (Keep It Simple Stupid) is a minimalist CLI note-taking application written in Go. It provides quick, friction-free note management from the terminal with notes stored in `~/.local/share/ks/` following the XDG Base Directory specification.

## Build and Development Commands

```bash
# Build the binary
go build

# Run the application
./ks [flags] [arguments]

# Install system-wide (optional)
sudo cp ks /usr/local/bin/

# Install for user only (optional)
mkdir -p ~/.local/bin && cp ks ~/.local/bin/
```

## Architecture

### Single-File Design

The entire application is contained in `main.go` (~1500 lines). This monolithic structure is intentional for simplicity and aligns with the minimalist philosophy.

### Core Components

**Theme System** (main.go:22-96): Centralized styling using lipgloss with consistent color scheme:
- Global `theme` variable provides Primary, Secondary, Accent, Error, Success, Warning, Muted styles
- Used throughout all TUI components and CLI output for visual consistency

**Keybinding System** (main.go:98-172): Centralized keybinding definitions using bubbles/key:
- `keyMap` struct defines all viewport navigation keys (scroll, jump, help, quit)
- Implements `help.KeyMap` interface for automatic help text generation
- Reusable across all TUI components

**CLI Flag System**: Uses Go's `flag` package with both short and long flags:
- Command flags: `-w`, `-l`, `-r`, `-d`, `-a`, `-s`, `-h`
- Options: `--sort`, `-i/--interactive`, `--force`
- Only one command flag can be used at a time
- Default behavior (no flags): launches interactive browse mode

**Interactive TUI Models**: Built with Charm Bracelet ecosystem:
- `confirmModel` (main.go:416-491): Yes/No confirmation dialogs with styled options
- `writeInputModel` (main.go:511-699): Two-stage input (filename → content) with validation
- `noteViewerModel` (main.go:805-903): Scrollable viewport for reading notes with help toggle
- `noteListModel` (main.go:905-993): Interactive list browser with navigation and selection
- All models use `tea.WithAltScreen()` for non-destructive fullscreen display

**List Item Types**:
- `noteInfo` (main.go:1117-1130): Implements `list.Item` interface for note browsing
- `searchResult` (main.go:1132-1143): Implements `list.Item` for search results with match location

**Filename Validation**: Security-focused validation prevents path traversal attacks:
- `validateFilename()`: Blocks path separators, parent references, hidden files, empty names
- `suggestFilename()`: Auto-fixes invalid filenames by replacing separators with dashes

**Note Operations**:
- `writeNote()`: Overwrites existing notes with themed success messages
- `appendNote()`: Smart append with automatic newline detection
- `listNotes(sortBy, interactive)`: Supports both simple and interactive modes
- `readNote()`: Launches viewport or falls back to simple print (TTY detection)
- `deleteNote()`: Interactive confirmation unless `--force` flag used
- `searchNotes(keyword, interactive)`: Supports both simple and interactive result browsing

### Data Flow

1. Flags parsed → Single command validated (or default to browse mode if no flags)
2. Arguments processed → Multiple input modes (direct args, stdin, interactive TUI)
3. Filename validated → Suggestions provided if invalid
4. Notes directory ensured (`~/.local/share/ks/`)
5. TTY detection (`isTTY()`) → Choose interactive vs. simple output
6. Operation executed → TUI launched or simple output displayed
7. For interactive modes: Model state transitions → Final result returned

### Interactive Mode Flow

**Browse Mode** (`ks` with no args):
1. Call `listNotes("name", true)`
2. Create `noteListModel` with all notes
3. User navigates and selects note
4. On Enter: Call `readNote(selected.name)` → launches viewport
5. Viewport allows scrolling, help toggle, quit

**List Mode** (`ks -l -i`):
1. Load and sort notes based on `--sort` flag
2. If interactive + TTY: Launch `noteListModel`
3. Else: Simple formatted list output

**Search Mode** (`ks -s keyword -i`):
1. Scan all notes for keyword matches
2. Build `[]searchResult` with match locations
3. If interactive + TTY: Launch list with search results
4. User selects result → `readNote()` opens selected note

## Development Patterns

**Error Handling**: Uses `os.Exit(1)` for fatal errors with themed error messages via `theme.Error.Render()`. All filesystem operations check errors explicitly.

**TTY Detection**: `isTTY()` checks if stdout is a terminal. Interactive modes fall back to simple output for pipes/redirects, ensuring composability with other CLI tools.

**Input Flexibility**: Commands support multiple modes:
- Direct: `ks -w note.txt "content"`
- Stdin: `echo "content" | ks -w note.txt`
- Interactive: `ks -w` (prompts for filename and content)

**TUI State Management**:
- Bubble Tea models use state integers for multi-stage workflows
- Models implement Init(), Update(), View() pattern
- Window size messages trigger responsive layout recalculation
- Alt-screen mode preserves terminal state

**Theming**: All output (TUI and CLI) uses the global `theme` for consistency. Success messages use `theme.Success`, errors use `theme.Error`, etc.

## Common Development Tasks

**Adding New Commands**:
1. Add flag variables (short and long forms)
2. Implement command function with `(sortBy, interactive bool)` pattern if applicable
3. Add flag count validation in main()
4. Update `printUsage()` with examples
5. Apply theme styles to output messages

**Adding New TUI Components**:
1. Create model struct with required state
2. Implement `Init() tea.Cmd`, `Update(tea.Msg) (tea.Model, tea.Cmd)`, `View() string`
3. Handle `tea.KeyMsg` for navigation and `tea.WindowSizeMsg` for responsiveness
4. Use themed styles for visual consistency
5. Launch with `tea.NewProgram(model, tea.WithAltScreen())`

**Important Constraints**:
- Avoid mixing `fmt.Scanf()` with Bubble Tea (compete for stdin)
- Always check `isTTY()` before launching interactive modes
- Use `tea.WithAltScreen()` to preserve terminal history
- Keep keybindings consistent with `keyMap` definitions
