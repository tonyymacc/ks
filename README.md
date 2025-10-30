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
ks -r note.txt        # Read with scrollable viewer
ks -d note.txt        # Delete note
```

## Interactive Features

### Main Menu (REPL)
Run `ks` to launch the menu:
- **Browse Notes** - Navigate all notes with preview panel
- **Search Notes** - Find notes by keyword
- **Create New Note** - Interactive note creation
- **Change Theme** - Select from 4 beautiful color schemes
- **Help** - View command reference
- **Quit** - Exit application

The menu loops continuously - perfect for extended note-taking sessions.

### List View Keybindings
The preview panel is visible by default, showing note content as you navigate.

- `↑/↓` or `j/k` - Navigate notes
- `/` - Filter/search
- `Enter` - Open note in full viewer
- `p` - Toggle preview panel
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

Simple, focused commands for quick operations. For full features, use the interactive REPL menu.

| Command | Description | Example |
|---------|-------------|---------|
| `-w, --write` | Create/overwrite note | `ks -w todo.txt "Buy milk"` |
| `-a, --append` | Append to note | `ks -a todo.txt "Walk dog"` |
| `-r, --read` | Read note in viewer | `ks -r todo.txt` |
| `-d, --delete` | Delete note | `ks -d old.txt` |
| `-h, --help` | Show help | `ks -h` |

**Tip:** Run `ks` without flags to access browse, search, sorting, and all interactive features!

## Themes

Choose from 4 beautiful color schemes via the main menu:

- **Purple (Default)** - Elegant purple/magenta tones
- **Ocean** - Cool blue/cyan palette
- **Forest** - Natural green hues
- **Sunset** - Warm orange/red colors

Change themes anytime from the main menu → "Change Theme". No code editing or rebuilding required!

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
- Split-view preview panel (visible by default)
- In-app theme selector (4 themes)
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
- More themes

## License

MIT - Educational project, free to use and modify.
