# ks - Keep It Simple Stupid

A minimalist, fast CLI note-taking application built with Go. Designed for developers who want quick, friction-free note management from the terminal.

## Features

- **Write notes** - Create and overwrite notes with a single command
- **Append to notes** - Add content to existing notes without overwriting
- **List notes** - View all notes with file sizes and timestamps
- **Flexible sorting** - Sort notes by name, date, or size
- **Read notes** - Display note contents quickly
- **Delete notes** - Remove notes when you no longer need them
- **Search** - Find notes by filename or content (case-insensitive)
- **Simple interface** - Clean flag-based commands with short and long options

## Installation

### Prerequisites

- Go 1.18 or higher

### Build from Source

```bash
# Clone the repository
cd ~/Projects/ks

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
ks
ks --help
```

## Storage Location

Notes are stored in:
```
~/.local/share/ks/
```

This follows the XDG Base Directory specification and keeps your notes organized in a standard location.

## Examples & Workflows

### Daily Journal

```bash
# Create today's entry
ks -w journal.txt "$(date): Started working on new feature"

# Add throughout the day
ks -a journal.txt $'\n- Completed code review'
ks -a journal.txt $'\n- Fixed bug in authentication'
```

### Meeting Notes

```bash
# Quick meeting notes
ks -w standup-2025-10-29.txt "Sprint planning discussion"

# Add action items later
ks -a standup-2025-10-29.txt $'\n\nAction Items:\n- Review PR #123\n- Update docs'
```

### Todo List Management

```bash
# Create todo list
ks -w todo.txt "1. Code review\n2. Write tests\n3. Update docs"

# Check what's on the list
ks -r todo.txt

# Add more tasks
ks -a todo.txt $'\n4. Deploy to staging'
```

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

### Multi-line Notes

For longer notes, consider using here documents:

```bash
ks -w long-note.txt "$(cat <<'EOF'
This is a longer note
with multiple lines
and proper formatting.
EOF
)"
```

### Quick Note Templates

Create aliases for common note types:

```bash
# Add to your ~/.bashrc or ~/.zshrc
alias meeting='ks -w meeting-$(date +%Y-%m-%d).txt'
alias idea='ks -a ideas.txt'
```

## Error Handling

The application provides clear error messages:

- **Missing arguments**: Shows correct usage
- **File not found**: Indicates which file couldn't be found
- **Multiple flags**: Prevents conflicting operations
- **Read errors**: Skips files that can't be read

## Development

### Project Structure

```
ks/
├── main.go          # Main application code
├── go.mod           # Go module definition
├── ks               # Compiled binary
└── README.md        # This file
```

### Building

```bash
go build
```

### Testing

Run the application with various commands to test functionality:

```bash
# Test write
./ks -w test.txt "test content"

# Test list
./ks -l

# Test search
./ks -s test

# Test cleanup
./ks -d test.txt
```

## Technical Details

- **Language**: Go 1.25+
- **File Storage**: `~/.local/share/ks/`
- **File Format**: Plain text files
- **Permissions**:
  - Notes directory: `0755` (rwxr-xr-x)
  - Note files: `0644` (rw-r--r--)

## Contributing

This is a learning project, but improvements are welcome:

1. Test the application thoroughly
2. Ensure all existing features still work
3. Add examples to the README
4. Follow Go conventions and best practices

## License

This project is created for educational purposes. Feel free to use and modify as needed.

## Author

Built as a learning project to explore Go fundamentals:
- CLI application development
- File I/O operations
- Flag parsing
- String manipulation
- Sorting and searching
- Error handling

## Roadmap

Future enhancements being considered:

- **Categories/Subdirectories** - Organize notes in folders
- **Export** - Export all notes to a single file
- **Configuration** - Customize storage location and behavior
- **Tags** - Tag-based organization system
- **Encryption** - Protect sensitive notes
- **REPL Mode** - Interactive TUI using Charm Bracelet libraries

---

**Keep It Simple Stupid** - Because note-taking shouldn't be complicated.
