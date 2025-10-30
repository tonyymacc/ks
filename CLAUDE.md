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

The entire application is contained in `main.go` (~1000 lines). This monolithic structure is intentional for simplicity and aligns with the minimalist philosophy.

### Core Components

**CLI Flag System**: Uses Go's `flag` package with both short (`-w`, `-l`, `-r`, `-d`, `-a`, `-s`, `-h`) and long (`--write`, `--list`, etc.) flags. Only one command flag can be used at a time.

**Interactive TUI Modes**: Built with the Charm Bracelet ecosystem (`bubbletea`, `bubbles`, `lipgloss`):
- `confirmModel` (main.go:331-430): Yes/No confirmation dialogs with keyboard navigation
- `writeInputModel` (main.go:433-573): Two-stage input (filename → content) with validation and suggestions
- Interactive modes support multiple input methods: fully interactive (no args), filename-only (prompts for content), or piped stdin

**Filename Validation**: Security-focused validation prevents path traversal attacks:
- `validateFilename()` (main.go:280-303): Blocks path separators (`/`, `\`), parent references (`..`), hidden files (leading `.`), and empty names
- `suggestFilename()` (main.go:575-593): Attempts to auto-fix invalid filenames by replacing separators with dashes

**Note Operations**:
- `writeNote()`: Overwrites existing notes
- `appendNote()`: Smart append with automatic newline detection (checks last byte) and confirmation for new files
- `listNotes()`: Sortable by name/date/size with human-readable formatting
- `readNote()`: Displays note with formatted header
- `deleteNote()`: Interactive confirmation unless `--force` flag used
- `searchNotes()`: Case-insensitive search across filenames and content

### Data Flow

1. Flags parsed → Single command validated
2. Arguments processed → Multiple input modes (direct args, stdin, interactive TUI)
3. Filename validated → Suggestions provided if invalid
4. Notes directory ensured (`~/.local/share/ks/`)
5. Operation executed with error handling

## Development Patterns

**Error Handling**: The application uses `os.Exit(1)` for fatal errors with user-friendly messages. All filesystem operations check for errors explicitly.

**Input Flexibility**: Most commands support three modes:
- Direct: `ks -w note.txt "content"`
- Stdin: `echo "content" | ks -w note.txt`
- Interactive: `ks -w` (prompts for both filename and content)

**TUI State Management**: Bubble Tea models use state integers to track progression through multi-stage workflows (e.g., 0=filename, 1=content, 2=done).

## Common Development Tasks

When modifying the codebase, note that adding new commands requires:
1. Adding flag variables and registering both short and long forms
2. Implementing the command function following existing patterns
3. Adding flag count validation in main()
4. Updating `printUsage()` with the new command

The Charm Bracelet libraries provide the TUI framework - avoid mixing traditional `fmt.Scanf()` with Bubble Tea models as they compete for stdin.
