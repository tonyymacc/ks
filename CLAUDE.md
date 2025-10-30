# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`ks` (Keep Simple Notes) is an interactive TUI note-taking application written in Go with the Charm Bracelet ecosystem. It features a full REPL interface with menu system, split-view preview panel, and comprehensive keyboard navigation. Notes are stored in `~/.local/share/ks/` following the XDG Base Directory specification.

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

The entire application is contained in `main.go` (~1800 lines). This monolithic structure is intentional for simplicity and maintainability.

### Core Components

**REPL Loop** (`runREPL()`): Main application entry point when no flags provided:
- Displays main menu with 5 options (Browse, Search, Create, Help, Quit)
- Loops continuously until user quits
- Each menu action returns to menu after completion
- Handles all interactive workflows

**Theme System**: Centralized styling using lipgloss:
- Global `theme` variable with Primary, Secondary, Accent, Error, Success, Warning, Muted styles
- Used throughout all TUI components and CLI output for consistency

**Keybinding System**: Centralized definitions using bubbles/key:
- `keyMap` struct defines viewport navigation (scroll, jump, help, quit)
- Implements `help.KeyMap` interface for automatic help generation
- List view has additional keybindings for actions (n/e/d/s/p)

**CLI Flag System**: Uses Go's `flag` package:
- Command flags: `-w`, `-l`, `-r`, `-d`, `-a`, `-s`, `-h`
- Options: `--sort`, `--force`
- Interactive mode automatic when TTY detected (no `-i` flag needed)
- Default behavior (no flags): launches REPL menu

**Interactive TUI Models**:
- `menuModel`: Main menu with cursor navigation and action selection
- `confirmModel`: Yes/No dialogs with styled options
- `writeInputModel`: Two-stage input (filename → content) with validation
- `noteViewerModel`: Scrollable viewport with help toggle
- `noteListModel`: Enhanced list with preview panel, sorting, and actions
  - Includes `viewport` for split-view preview
  - Tracks `sortMode`, `action`, `showPreview` state
  - Handles n/e/d/s/p keybindings
  - Returns action type to caller for processing

**List Item Types**:
- `noteInfo`: Implements `list.Item` for note browsing
- `searchResult`: Implements `list.Item` for search results

**Action Handling**:
- `handleListAction()`: Processes actions from noteListModel (open/create/rename/delete/quit)
- `runInteractiveCreate()`: Launches note creation workflow
- `runInteractiveRename()`: Prompts for new filename and renames
- `runInteractiveSearch()`: Prompts for keyword and launches search
- All actions return to REPL menu when complete

**Note Operations**:
- `writeNote()`: Overwrites with themed messages
- `appendNote()`: Smart append with newline detection
- `listNotes(sortBy, interactive)`: Simple or interactive mode
- `readNote()`: Viewport with TTY fallback
- `deleteNote()`: Confirmation unless --force
- `searchNotes(keyword, interactive)`: Simple or interactive results
- `sortNotes()`: Helper for dynamic re-sorting

### Data Flow

**REPL Mode** (no flags):
1. Check TTY → Launch `runREPL()` or show help
2. Display `menuModel` → Wait for user selection
3. Execute action (browse/search/create/help)
4. Action completes → Return to menu (loop)
5. User selects Quit → Exit application

**CLI Mode** (with flags):
1. Parse flags → Validate single command
2. Process arguments → Handle stdin/direct/interactive input
3. Validate filenames → Provide suggestions if invalid
4. Ensure notes directory exists
5. TTY detection → Choose interactive vs. simple mode
6. Execute operation → Return to shell

### REPL Workflow

**Main Menu Loop**:
```
┌─────────────────┐
│  Display Menu   │
└────────┬────────┘
         │
    ┌────▼─────┐
    │  Choice  │
    └────┬─────┘
         │
    ┌────▼────────────────────────────┐
    │ Browse │ Search │ Create │ Help │ → Quit (exit)
    └────┬────────┬────────┬─────┬────┘
         │        │        │     │
         ▼        ▼        ▼     ▼
    listNotes  search  create  help
         │        │        │     │
         └────────┴────────┴─────┘
                  │
             Back to Menu
```

**List View Actions** (from noteListModel):
- Press `Enter` → `action="open"` → `readNote()` → viewport → back to menu
- Press `n` → `action="create"` → `runInteractiveCreate()` → back to menu
- Press `e` → `action="rename"` → `runInteractiveRename()` → back to menu
- Press `d` → `action="delete"` → `deleteNote()` → back to menu
- Press `s` → Re-sort in-place (stay in list)
- Press `p` → Toggle preview (stay in list)
- Press `q`/`Esc` → Back to menu

**Preview Panel**:
- Triggered by `p` key in list view
- Split screen: list (left 50%) | preview (right 50%)
- Preview updates automatically when navigating
- Viewport allows scrolling through preview content

## Development Patterns

**Error Handling**: Uses `os.Exit(1)` for fatal errors with themed error messages via `theme.Error.Render()`. All filesystem operations check errors explicitly.

**TTY Detection**: `isTTY()` checks if stdout is a terminal. Interactive modes fall back to simple output for pipes/redirects, ensuring composability with other CLI tools.

**REPL Pattern**:
- Main loop in `runREPL()` displays menu → executes action → returns to menu
- Actions communicate via model state (`action`, `selected`)
- `handleListAction()` dispatches to appropriate handlers
- All interactive functions return to REPL, never exit directly

**Action Passing**:
- TUI models set `action` field before quitting
- Caller checks final model state and dispatches
- Actions: "open", "create", "rename", "delete", "quit"
- Allows same model to trigger different workflows

**Input Flexibility**: Commands support multiple modes:
- Direct: `ks -w note.txt "content"`
- Stdin: `echo "content" | ks -w note.txt`
- Interactive: `ks -w` (prompts for filename and content)

**TUI State Management**:
- Bubble Tea models use state fields for multi-stage workflows
- Models implement Init(), Update(), View() pattern
- Window size messages trigger responsive layout recalculation
- Alt-screen mode preserves terminal state
- Preview panel uses `showPreview` boolean to toggle split view

**Theming**: All output (TUI and CLI) uses the global `theme` for consistency. Success messages use `theme.Success`, errors use `theme.Error`, etc.

## Common Development Tasks

**Adding New CLI Commands**:
1. Add flag variables (short and long forms)
2. Implement command function with TTY auto-detection
3. Add flag count validation in main()
4. Update `printUsage()` with examples
5. Apply theme styles to output messages

**Adding REPL Menu Actions**:
1. Add menu choice to `menuModel.choices` array
2. Add case in `runREPL()` switch statement
3. Create `runInteractive<Action>()` function
4. Function must return to caller (not exit)
5. Test menu loop returns properly

**Adding List Keybindings**:
1. Add keybinding to `noteListModel.Update()` switch
2. Set appropriate `action` field value
3. Add handler case in `handleListAction()`
4. Update `AdditionalShortHelpKeys` in `newNoteListModel()`
5. Test action returns to menu

**Adding New TUI Components**:
1. Create model struct with required state fields
2. Implement `Init() tea.Cmd`, `Update(tea.Msg) (tea.Model, tea.Cmd)`, `View() string`
3. Handle `tea.KeyMsg` for navigation and `tea.WindowSizeMsg` for responsiveness
4. Use themed styles for visual consistency
5. Launch with `tea.NewProgram(model, tea.WithAltScreen())`
6. If used in REPL, ensure it returns to menu

**Important Constraints**:
- Never call `os.Exit()` from REPL-launched functions
- Avoid mixing `fmt.Scanf()` with Bubble Tea (compete for stdin)
- Always check `isTTY()` before launching interactive modes
- Use `tea.WithAltScreen()` to preserve terminal history
- Keep keybindings consistent across views
- Update help text when adding keybindings
