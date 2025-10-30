# ks - Keep Simple Notes

A beautiful, interactive note-taking TUI built with Go and [Charm Bracelet](https://charm.sh/). Manage notes from your terminal with a rich, keyboard-driven interface.

## Installation

**Prerequisites:** Go 1.18+

```bash
git clone https://github.com/tonyymacc/ks.git
cd ks
go build
sudo cp ks /usr/local/bin/  # or: cp ks ~/.local/bin/
```

## Quick Start

```bash
ks                    # Launch interactive menu (REPL mode)
ks -w note.txt "..."  # Quick write
ks -l                 # Interactive list (with TTY)
ks -r note.txt        # Read with scrollable viewer
ks -s keyword         # Interactive search
```

## Interactive Features

### Main Menu (REPL)
Run `ks` to launch the menu:
- **Browse Notes** - Navigate all notes with preview
- **Search Notes** - Find notes by keyword
- **Create New Note** - Interactive note creation
- **Help** - View command reference
- **Quit** - Exit application

The menu loops continuously - perfect for extended note-taking sessions.

### List View Keybindings
- `↑/↓` or `j/k` - Navigate
- `/` - Filter/search
- `Enter` - Open note
- `p` - Toggle preview panel (split view)
- `s` - Cycle sort (name → date → size)
- `n` - Create new note
- `e` - Rename selected note
- `d` - Delete note
- `q` - Back to menu

### Note Viewer
- `↑/↓` or `j/k` - Scroll line by line
- `u/d` or `Ctrl+U/D` - Half-page scroll
- `f/b` or `PgDn/PgUp` - Full-page scroll
- `g/G` - Jump to top/bottom
- `?` - Toggle help
- `q` - Close viewer

## CLI Commands

All commands support both short and long forms. Interactive mode activates automatically when using a TTY.

| Command | Description | Example |
|---------|-------------|---------|
| `-w, --write` | Create/overwrite note | `ks -w todo.txt "Buy milk"` |
| `-a, --append` | Append to note | `ks -a todo.txt "Walk dog"` |
| `-l, --list` | List all notes | `ks -l --sort date` |
| `-r, --read` | Read note in viewer | `ks -r todo.txt` |
| `-d, --delete` | Delete note | `ks -d old.txt` |
| `-s, --search` | Search notes | `ks -s golang` |
| `-h, --help` | Show help | `ks -h` |

**Sort options:** `--sort name` (default), `--sort date`, `--sort size`

## Theming & Customization

The application uses a built-in color theme with lipgloss. Colors are defined in `main.go`:

```go
theme.Primary    // Purple/magenta for headers and important text
theme.Secondary  // Gray for metadata and descriptions
theme.Accent     // Pink for highlights and matches
theme.Error      // Red for errors
theme.Success    // Green for success messages
theme.Warning    // Orange for warnings
theme.Muted      // Dim gray for help text
```

To customize colors, edit the `defaultTheme()` function in `main.go` and rebuild.

## Storage

Notes are stored in `~/.local/share/ks/` (XDG Base Directory specification).

## Tips

**Newlines in bash:** Use `$'\n'` for actual newlines:
```bash
ks -w note.txt $'Line 1\nLine 2'  # Correct
ks -w note.txt "Line 1\nLine 2"    # Wrong (literal \n)
```

**Piping:** Works seamlessly with pipes and redirects:
```bash
echo "content" | ks -w note.txt    # Write from stdin
ks -r note.txt | grep "keyword"    # Pipe note content
```

## Roadmap

✅ Completed:
- Interactive REPL with main menu
- Split-view preview panel
- Scrollable viewer with help toggle
- In-app note creation/renaming/deletion
- Dynamic sorting (name/date/size)
- Comprehensive keybindings

🔮 Future:
- Categories/subdirectories
- Tags system
- Export all notes
- Configuration file
- Editor integration ($EDITOR)
- Encryption

## License

MIT - Educational project, free to use and modify.
