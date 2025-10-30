# ks

A minimalist, fast CLI note-taking application built with Go. Designed for developers who want quick, friction-free note management from the terminal.

## Installation

### Prerequisites

- Go 1.18 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/tonyymacc/ks.git

# Build the binary
go build

# (Optional) Install system-wide
sudo cp ks /usr/local/bin/
# or for user-only install
mkdir -p ~/.local/bin
cp ks ~/.local/bin/
```

## Usage

### Quick Start

```bash
# Create a note
ks -w mynote.txt "This is my first note"

# List all notes
ks -l

# Read a note
ks -r mynote.txt

# Search for notes
ks -s "keyword"

# Delete a note
ks -d mynote.txt
```

### Commands

#### Write a Note

Create a new note or overwrite an existing one.

```bash
ks -w <filename> <content>
ks --write <filename> <content>
```

**Examples:**
```bash
ks -w todo.txt "Buy groceries"
ks --write meeting-notes.txt "Discussed Q1 goals"
```

#### Append to a Note

Add content to the end of an existing note. Creates the note if it doesn't exist.

```bash
ks -a <filename> <content>
ks --append <filename> <content>
```

**Examples:**
```bash
# Use $'\n' for actual newlines in bash
ks -a todo.txt $'\nFinish project'
ks --append ideas.txt " - Another idea"
```

#### List Notes

Display all notes with file sizes, timestamps, and optional sorting.

```bash
ks -l [--sort <order>]
ks --list [--sort <order>]
```

**Sort Options:**
- `name` - Alphabetical by filename (default)
- `date` - By modification time (newest first)
- `size` - By file size (largest first)

**Examples:**
```bash
ks -l                  # List with default sort (by name)
ks -l --sort date      # List newest notes first
ks -l --sort size      # List largest notes first
```

**Output:**
```
Notes:
  - golang-notes.txt                   69 B  (modified: 2025-10-29 14:10)
  - meeting-notes.txt                 2.1 KB  (modified: 2025-10-29 15:30)
  - todo.txt                           45 B  (modified: 2025-10-29 16:00)
```

#### Read a Note

Display the contents of a note.

```bash
ks -r <filename>
ks --read <filename>
```

**Examples:**
```bash
ks -r todo.txt
ks --read meeting-notes.txt
```

**Output:**
```
=== todo.txt ===
Buy groceries
Finish project
```

#### Delete a Note

Permanently remove a note.

```bash
ks -d <filename>
ks --delete <filename>
```

**Examples:**
```bash
ks -d old-note.txt
ks --delete temp.txt
```

#### Search Notes

Search for keywords in note filenames and content. Case-insensitive.

```bash
ks -s <keyword>
ks --search <keyword>
```

**Examples:**
```bash
ks -s golang          # Find notes about golang
ks --search meeting   # Find meeting notes
```

**Output:**
```
Searching for: golang

  - golang-notes.txt               (match in: filename and content)
  - tutorial.txt                   (match in: content)

Found 2 match(es).
```

### Help

Display usage information and examples.

```bash
ks -h
ks --help
```

## Storage Location

Notes are stored in:
```
~/.local/share/ks/
```

This follows the XDG Base Directory specification and keeps your notes organized in a standard location.

### Finding Notes

```bash
# Find all golang-related notes
ks -s golang

# Find notes modified recently
ks -l --sort date

# Find large notes that might need cleanup
ks -l --sort size
```

## Tips & Tricks

### Using Newlines

In bash, use `$'\n'` for actual newlines:

```bash
# Wrong - will show literal \n
ks -w note.txt "Line 1\nLine 2"

# Correct - actual newline
ks -w note.txt $'Line 1\nLine 2'
```

### Quick Note Templates

Create aliases for common note types:

```bash
# Add to your ~/.bashrc or ~/.zshrc
alias meeting='ks -w meeting-$(date +%Y-%m-%d).txt'
alias idea='ks -a ideas.txt'
```

## Contributing

This is a learning project, but improvements are welcome:

1. Test the application thoroughly
2. Ensure all existing features still work
3. Add examples to the README
4. Follow Go conventions and best practices

## License

This project is created for educational purposes. Feel free to use and modify as needed.

## Author

Built as a learning project to explore Go fundamentals

## Roadmap

Future enhancements being considered:

- **Categories/Subdirectories** - Organize notes in folders
- **Export** - Export all notes to a single file
- **Configuration** - Customize storage location and behavior
- **Tags** - Tag-based organization system
- **Encryption** - Protect sensitive notes
- **REPL Mode** - Interactive TUI using Charm Bracelet libraries
