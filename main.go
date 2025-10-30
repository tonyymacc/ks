package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// Theme defines the color scheme and styles for the application
type Theme struct {
	Primary       lipgloss.Style
	Secondary     lipgloss.Style
	Accent        lipgloss.Style
	Error         lipgloss.Style
	Success       lipgloss.Style
	Warning       lipgloss.Style
	Muted         lipgloss.Style
	Border        lipgloss.Style
	Header        lipgloss.Style
	Highlight     lipgloss.Style
	Selected      lipgloss.Style
	Unselected    lipgloss.Style
}

// defaultTheme creates the default purple theme
func defaultTheme() Theme {
	return Theme{
		Primary: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("170")),

		Secondary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")),

		Accent: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("213")),

		Error: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196")),

		Success: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42")),

		Warning: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")),

		Muted: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("170")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Highlight: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("235")),

		Selected: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("170")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Unselected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1),
	}
}

// oceanTheme creates a blue/cyan ocean theme
func oceanTheme() Theme {
	return Theme{
		Primary: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")),

		Secondary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")),

		Accent: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("51")),

		Error: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196")),

		Success: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42")),

		Warning: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")),

		Muted: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("39")).
			Padding(0, 1),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("51")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Highlight: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("235")),

		Selected: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Unselected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1),
	}
}

// forestTheme creates a green forest theme
func forestTheme() Theme {
	return Theme{
		Primary: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("34")),

		Secondary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")),

		Accent: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("46")),

		Error: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196")),

		Success: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42")),

		Warning: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")),

		Muted: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("34")).
			Padding(0, 1),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("46")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Highlight: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("235")),

		Selected: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("34")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Unselected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1),
	}
}

// sunsetTheme creates an orange/red sunset theme
func sunsetTheme() Theme {
	return Theme{
		Primary: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("208")),

		Secondary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")),

		Accent: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")),

		Error: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("196")),

		Success: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42")),

		Warning: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")),

		Muted: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("208")).
			Padding(0, 1),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("214")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Highlight: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("235")),

		Selected: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("208")).
			Background(lipgloss.Color("235")).
			Padding(0, 1),

		Unselected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1),
	}
}

// Global theme instance
var theme = defaultTheme()
var currentThemeName = "Purple (Default)"

// Available themes map
var themes = map[string]func() Theme{
	"Purple (Default)": defaultTheme,
	"Ocean":            oceanTheme,
	"Forest":           forestTheme,
	"Sunset":           sunsetTheme,
}

// keyMap defines keybindings for the viewport
type keyMap struct {
	Up       key.Binding
	Down     key.Binding
	PageUp   key.Binding
	PageDown key.Binding
	HalfUp   key.Binding
	HalfDown key.Binding
	Top      key.Binding
	Bottom   key.Binding
	Quit     key.Binding
	Help     key.Binding
}

// ShortHelp returns a quick one-line help
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns detailed help with sections
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.PageUp, k.PageDown},
		{k.HalfUp, k.HalfDown, k.Top, k.Bottom},
		{k.Help, k.Quit},
	}
}

// defaultKeyMap creates the default keybindings
func defaultKeyMap() keyMap {
	return keyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "scroll up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "scroll down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup", "b"),
			key.WithHelp("pgup/b", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("pgdown", "f", " "),
			key.WithHelp("pgdn/f/space", "page down"),
		),
		HalfUp: key.NewBinding(
			key.WithKeys("u", "ctrl+u"),
			key.WithHelp("u", "half page up"),
		),
		HalfDown: key.NewBinding(
			key.WithKeys("d", "ctrl+d"),
			key.WithHelp("d", "half page down"),
		),
		Top: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to top"),
		),
		Bottom: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to bottom"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

var keys = defaultKeyMap()

// runREPL starts the interactive REPL mode
func runREPL() {
	for {
		// Show main menu
		m := newMenuModel()
		p := tea.NewProgram(m, tea.WithAltScreen())
		result, err := p.Run()
		if err != nil {
			fmt.Println(theme.Error.Render("✗ Error: " + err.Error()))
			os.Exit(1)
		}

		menu := result.(menuModel)

		switch menu.selected {
		case "Browse Notes":
			listNotes("name", true)
		case "Search Notes":
			runInteractiveSearch()
		case "Create New Note":
			runInteractiveCreate()
		case "Change Theme":
			runThemeSelector()
		case "Help":
			printUsage()
			fmt.Println("\n" + theme.Muted.Render("Press Enter to continue..."))
			fmt.Scanln()
		case "Quit", "quit":
			fmt.Println(theme.Success.Render("Goodbye!"))
			return
		}
	}
}

// runInteractiveSearch prompts for a search keyword and runs interactive search
func runInteractiveSearch() {
	ti := textinput.New()
	ti.Placeholder = "Enter search keyword..."
	ti.Focus()

	fmt.Print(theme.Primary.Render("Search: "))

	// Simple input for now - could make this a full TUI
	var keyword string
	fmt.Scanln(&keyword)

	if keyword != "" {
		searchNotes(keyword, true)
	}
}

// runInteractiveCreate launches the interactive note creation
func runInteractiveCreate() {
	filename, content, ok := interactiveWrite()
	if ok {
		writeNote(filename, content)
	}
}

func main() {
	// Define flag variables - keep it simple
	var writeFlag, readFlag, deleteFlag, appendFlag, helpFlag bool

	// Register short flags
	flag.BoolVar(&writeFlag, "w", false, "Write a note")
	flag.BoolVar(&readFlag, "r", false, "Read a note")
	flag.BoolVar(&deleteFlag, "d", false, "Delete a note")
	flag.BoolVar(&appendFlag, "a", false, "Append to a note")
	flag.BoolVar(&helpFlag, "h", false, "Show help message")

	// Register long flags
	flag.BoolVar(&writeFlag, "write", false, "Write a note")
	flag.BoolVar(&readFlag, "read", false, "Read a note")
	flag.BoolVar(&deleteFlag, "delete", false, "Delete a note")
	flag.BoolVar(&appendFlag, "append", false, "Append to a note")
	flag.BoolVar(&helpFlag, "help", false, "Show help message")

	// Force flag (skip confirmations)
	var forceFlag bool
	flag.BoolVar(&forceFlag, "force", false, "Skip confirmation prompts")

	// Custom usage message
	flag.Usage = printUsage

	// Parse the flags
	flag.Parse()

	// Handle help flag explicitly
	if helpFlag {
		printUsage()
		os.Exit(0)
	}

	// Get remaining arguments after flags
	args := flag.Args()

	// If no flags provided, launch REPL mode
	if flag.NFlag() == 0 {
		if isTTY() {
			runREPL()
		} else {
			printUsage()
		}
		return
	}

	// Check that only one flag is used at a time
	flagCount := 0
	if writeFlag {
		flagCount++
	}
	if readFlag {
		flagCount++
	}
	if deleteFlag {
		flagCount++
	}
	if appendFlag {
		flagCount++
	}

	if flagCount > 1 {
		fmt.Println("Error: Only one command flag can be used at a time")
		printUsage()
		os.Exit(1)
	}

	// Execute the appropriate command based on flag
	if writeFlag {
		var filename, note string

		// Multiple modes: 0 args (interactive), 1 arg (stdin or interactive content), 2 args (direct)
		if len(args) == 0 {
			// Interactive mode - prompt for filename and content
			filename, note, ok := interactiveWrite()
			if !ok {
				fmt.Println("Write cancelled.")
				os.Exit(1)
			}
			writeNote(filename, note)
		} else if len(args) == 1 {
			// Check if stdin has content
			stdinContent, hasStdin := readFromStdin()
			if hasStdin {
				// stdin mode - use provided filename and stdin content
				filename = args[0]
				note = stdinContent
				writeNote(filename, note)
			} else {
				// Interactive content mode - use provided filename, prompt for content
				filename = args[0]
				// Validate filename first
				if err := validateFilename(filename); err != nil {
					suggested := suggestFilename(filename)
					if suggested != "" {
						fmt.Printf("Invalid filename: %v\nSuggestion: %s\n", err, suggested)
					} else {
						fmt.Printf("Invalid filename: %v\n", err)
					}
					os.Exit(1)
				}
				note, ok := interactiveContent(filename)
				if !ok {
					fmt.Println("Write cancelled.")
					os.Exit(1)
				}
				writeNote(filename, note)
			}
		} else if len(args) == 2 {
			filename = args[0]
			note = args[1]
			writeNote(filename, note)
		} else {
			fmt.Println("Usage: ks -w <filename> <note>")
			fmt.Println("   or: ks --write <filename> <note>")
			fmt.Println("   or: echo \"content\" | ks -w <filename>")
			fmt.Println("   or: ks -w <filename> (interactive content)")
			fmt.Println("   or: ks -w (fully interactive)")
			os.Exit(1)
		}
	} else if readFlag {
		if len(args) != 1 {
			fmt.Println("Usage: ks -r <filename>")
			fmt.Println("   or: ks --read <filename>")
			os.Exit(1)
		}
		readNote(args[0])
	} else if deleteFlag {
		if len(args) != 1 {
			fmt.Println("Usage: ks -d <filename>")
			fmt.Println("   or: ks --delete <filename>")
			os.Exit(1)
		}
		deleteNote(args[0], forceFlag)
	} else if appendFlag {
		var filename, note string

		// Multiple modes: 0 args (interactive), 1 arg (stdin or interactive content), 2 args (direct)
		if len(args) == 0 {
			// Interactive mode - prompt for filename and content
			filename, note, ok := interactiveAppend()
			if !ok {
				fmt.Println("Append cancelled.")
				os.Exit(1)
			}
			appendNote(filename, note)
		} else if len(args) == 1 {
			// Check if stdin has content
			stdinContent, hasStdin := readFromStdin()
			if hasStdin {
				// stdin mode - use provided filename and stdin content
				filename = args[0]
				note = stdinContent
				appendNote(filename, note)
			} else {
				// Interactive content mode - use provided filename, prompt for content
				filename = args[0]
				// Validate filename first
				if err := validateFilename(filename); err != nil {
					suggested := suggestFilename(filename)
					if suggested != "" {
						fmt.Printf("Invalid filename: %v\nSuggestion: %s\n", err, suggested)
					} else {
						fmt.Printf("Invalid filename: %v\n", err)
					}
					os.Exit(1)
				}
				note, ok := interactiveContent(filename)
				if !ok {
					fmt.Println("Append cancelled.")
					os.Exit(1)
				}
				appendNote(filename, note)
			}
		} else if len(args) == 2 {
			filename = args[0]
			note = args[1]
			appendNote(filename, note)
		} else {
			fmt.Println("Usage: ks -a <filename> <note>")
			fmt.Println("   or: ks --append <filename> <note>")
			fmt.Println("   or: echo \"content\" | ks -a <filename>")
			fmt.Println("   or: ks -a <filename> (interactive content)")
			fmt.Println("   or: ks -a (fully interactive)")
			os.Exit(1)
		}
	}
}

// printUsage displays the help message
func printUsage() {
	fmt.Println(theme.Header.Render(" ks - Keep Simple Notes "))
	fmt.Println("\nUsage:")
	fmt.Println("  ks                                Launch interactive REPL menu")
	fmt.Println("  ks [flags] [arguments]            Run specific command")
	fmt.Println("\nFlags:")
	fmt.Println("  -w, --write <filename> <note>    Write a note")
	fmt.Println("  -a, --append <filename> <note>   Append to a note")
	fmt.Println("  -r, --read <filename>            Read a note")
	fmt.Println("  -d, --delete <filename>          Delete a note")
	fmt.Println("  -h, --help                       Show this help")
	fmt.Println("\nExamples:")
	fmt.Println("  ks                                # Launch REPL menu")
	fmt.Println("  ks -w note.txt \"My note\"          # Quick write")
	fmt.Println("  ks -a note.txt \"More content\"     # Quick append")
	fmt.Println("  ks -r note.txt                    # Read note")
	fmt.Println("  ks -d note.txt                    # Delete note")
	fmt.Println("\nTip: Run 'ks' without flags to access all features interactively!")
}

// getNotesDir returns the path to the notes directory
func getNotesDir() (string, error) {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Build the notes directory path: ~/.local/share/ks
	notesDir := filepath.Join(home, ".local", "share", "ks")

	// Create the directory if it doesn't exist
	err = os.MkdirAll(notesDir, 0755)
	if err != nil {
		return "", err
	}

	return notesDir, nil
}

// validateFilename ensures the filename is safe and doesn't contain path traversal attempts
func validateFilename(filename string) error {
	// Check for empty filename
	if filename == "" {
		return fmt.Errorf("filename cannot be empty")
	}

	// Check for path separators (prevents directory traversal)
	if strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		return fmt.Errorf("filename cannot contain path separators (/ or \\)")
	}

	// Check for parent directory references
	if strings.Contains(filename, "..") {
		return fmt.Errorf("filename cannot contain '..'")
	}

	// Check if filename starts with a dot (hidden files - optional security measure)
	if strings.HasPrefix(filename, ".") {
		return fmt.Errorf("filename cannot start with '.' (hidden files not allowed)")
	}

	return nil
}

// readFromStdin reads content from standard input
// Returns the content and true if stdin has data, or empty string and false if not
func readFromStdin() (string, bool) {
	// Check if stdin is a pipe or redirect (not a terminal)
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", false
	}

	// Check if stdin is a pipe or regular file (has data)
	// ModeCharDevice means it's an interactive terminal (no piped data)
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return "", false
	}

	// Read all data from stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading from stdin: %v\n", err)
		return "", false
	}

	return string(data), true
}

// isTTY checks if stdout is connected to a terminal
func isTTY() bool {
	stat, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}

// confirmModel is a Bubble Tea model for yes/no confirmation
type confirmModel struct {
	question string
	answer   bool
	cursor   int // 0 = No, 1 = Yes
	quitting bool
}

// Legacy style aliases (for backwards compatibility during transition)
var (
	selectedStyle   = theme.Selected
	unselectedStyle = theme.Unselected
	boxStyle        = theme.Border
)

func (m confirmModel) Init() tea.Cmd {
	return nil
}

func (m confirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			m.cursor = 0 // No
		case "right", "l":
			m.cursor = 1 // Yes
		case "y", "Y":
			m.answer = true
			m.quitting = true
			return m, tea.Quit
		case "n", "N":
			m.answer = false
			m.quitting = true
			return m, tea.Quit
		case "enter":
			m.answer = (m.cursor == 1)
			m.quitting = true
			return m, tea.Quit
		case "q", "esc", "ctrl+c":
			m.answer = false
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m confirmModel) View() string {
	if m.quitting {
		return ""
	}

	// Build the options with cursor
	var noOption, yesOption string
	if m.cursor == 0 {
		noOption = selectedStyle.Render("No")
		yesOption = unselectedStyle.Render("Yes")
	} else {
		noOption = unselectedStyle.Render("No")
		yesOption = selectedStyle.Render("Yes")
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, noOption, "  ", yesOption)
}

// confirm shows a yes/no prompt and returns true if user confirms
func confirm(question string) bool {
	// Print the question before showing the prompt
	fmt.Println(question)

	m := confirmModel{
		question: question,
		answer:   false,
		cursor:   0, // Default to "No" for safety
		quitting: false,
	}

	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running confirmation: %v\n", err)
		return false
	}

	if final, ok := finalModel.(confirmModel); ok {
		return final.answer
	}
	return false
}

// writeInputModel handles interactive write mode
type writeInputModel struct {
	state         int // 0 = filename input, 1 = content input, 2 = done
	filenameInput textinput.Model
	contentInput  textarea.Model
	filename      string
	content       string
	validationErr string
	quitting      bool
}

func newWriteInputModel() writeInputModel {
	ti := textinput.New()
	ti.Placeholder = "note.txt"
	ti.Focus()
	ti.CharLimit = 255
	ti.Width = 50

	ta := textarea.New()
	ta.Placeholder = "Write your note here..."
	ta.ShowLineNumbers = false
	ta.CharLimit = 0

	return writeInputModel{
		state:         0,
		filenameInput: ti,
		contentInput:  ta,
		validationErr: "",
		quitting:      false,
	}
}

func (m writeInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m writeInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			if m.state == 1 {
				// Ctrl+C or Esc in content mode asks for confirmation
				return m, nil
			}
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if m.state == 0 {
				// Validate filename
				filename := strings.TrimSpace(m.filenameInput.Value())
				if filename == "" {
					m.validationErr = "Filename cannot be empty"
					return m, nil
				}

				if err := validateFilename(filename); err != nil {
					// Check if we can suggest a fix
					suggested := suggestFilename(filename)
					if suggested != "" && suggested != filename {
						m.validationErr = fmt.Sprintf("%v - Suggestion: %s (press Tab to use)", err, suggested)
					} else {
						m.validationErr = err.Error()
					}
					return m, nil
				}

				// Filename is valid, move to content input
				m.filename = filename
				m.validationErr = ""
				m.state = 1
				m.contentInput.Focus()
				return m, textarea.Blink

			}

		case "tab":
			if m.state == 0 && m.validationErr != "" {
				// Apply suggested filename if available
				filename := m.filenameInput.Value()
				suggested := suggestFilename(filename)
				if suggested != "" {
					m.filenameInput.SetValue(suggested)
					m.validationErr = ""
				}
			}

		case "ctrl+d":
			if m.state == 1 {
				// Save and quit
				m.content = m.contentInput.Value()
				m.state = 2
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	// Update the active input
	if m.state == 0 {
		m.filenameInput, cmd = m.filenameInput.Update(msg)
	} else if m.state == 1 {
		m.contentInput, cmd = m.contentInput.Update(msg)
	}

	return m, cmd
}

func (m writeInputModel) View() string {
	if m.quitting {
		return ""
	}

	var s string

	if m.state == 0 {
		// Filename input stage
		s = theme.Primary.Render("Enter filename:") + "\n"
		s += m.filenameInput.View() + "\n"

		if m.validationErr != "" {
			s += "\n" + theme.Error.Render("✗ "+m.validationErr) + "\n"
		}

		s += "\n" + theme.Muted.Render("Enter to continue • Esc to cancel")

	} else if m.state == 1 {
		// Content input stage
		s = theme.Primary.Render("Writing to: ") + theme.Accent.Render(m.filename) + "\n\n"
		s += m.contentInput.View() + "\n\n"

		s += theme.Muted.Render("Ctrl+D to save • Esc to cancel")
	}

	return s
}

// suggestFilename attempts to fix common filename issues
func suggestFilename(filename string) string {
	// Replace path separators with dashes
	suggested := strings.ReplaceAll(filename, "/", "-")
	suggested = strings.ReplaceAll(suggested, "\\", "-")

	// Remove leading dots
	suggested = strings.TrimPrefix(suggested, ".")

	// Remove parent directory references
	suggested = strings.ReplaceAll(suggested, "..", "")

	// If we made changes, return the suggestion
	if suggested != filename && suggested != "" {
		return suggested
	}

	return ""
}

// interactiveWrite launches the interactive write mode
func interactiveWrite() (string, string, bool) {
	m := newWriteInputModel()
	p := tea.NewProgram(m)
	finalModel, err := p.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", "", false
	}

	if final, ok := finalModel.(writeInputModel); ok {
		if final.state == 2 {
			return final.filename, final.content, true
		}
	}

	return "", "", false
}

// interactiveAppend launches the interactive append mode (same as write but different prompt)
func interactiveAppend() (string, string, bool) {
	return interactiveWrite() // Same logic for now
}

// interactiveContent prompts only for content (filename already provided)
func interactiveContent(filename string) (string, bool) {
	ta := textarea.New()
	ta.Placeholder = "Write your note here..."
	ta.ShowLineNumbers = false
	ta.CharLimit = 0
	ta.Focus()

	m := writeInputModel{
		state:         1, // Skip directly to content input
		filenameInput: textinput.Model{},
		contentInput:  ta,
		filename:      filename,
		content:       "",
		validationErr: "",
		quitting:      false,
	}

	p := tea.NewProgram(m)
	finalModel, err := p.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", false
	}

	if final, ok := finalModel.(writeInputModel); ok {
		if final.state == 2 {
			return final.content, true
		}
	}

	return "", false
}

// noteViewerModel is a viewport-based note reader
type noteViewerModel struct {
	viewport viewport.Model
	help     help.Model
	filename string
	content  string
	ready    bool
	showHelp bool
	quitting bool
}

func newNoteViewerModel(filename, content string) noteViewerModel {
	h := help.New()
	h.ShowAll = false // Start with short help

	return noteViewerModel{
		filename: filename,
		content:  content,
		help:     h,
		ready:    false,
		showHelp: false,
		quitting: false,
	}
}

func (m noteViewerModel) Init() tea.Cmd {
	return nil
}

func (m noteViewerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Check for quit
		if key.Matches(msg, keys.Quit) {
			m.quitting = true
			return m, tea.Quit
		}

		// Check for help toggle
		if key.Matches(msg, keys.Help) {
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}

	case tea.WindowSizeMsg:
		headerHeight := 3
		footerHeight := 3
		if m.help.ShowAll {
			footerHeight = 6 // More space for expanded help
		}
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Initialize viewport with terminal size
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		m.help.Width = msg.Width
	}

	// Handle viewport scrolling
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m noteViewerModel) View() string {
	if m.quitting {
		return ""
	}

	if !m.ready {
		return "\n  Initializing..."
	}

	// Create header
	header := theme.Header.Render(" " + m.filename + " ")

	// Create footer with scroll position and help
	scrollPercent := int(m.viewport.ScrollPercent() * 100)
	scrollInfo := theme.Secondary.Render(fmt.Sprintf(" %d%% ", scrollPercent))

	// Render help using the help component
	helpView := m.help.View(keys)
	helpStyled := theme.Muted.Render(helpView)

	footer := lipgloss.JoinHorizontal(lipgloss.Top, scrollInfo, helpStyled)

	// Combine header, viewport content, and footer
	return fmt.Sprintf("%s\n%s\n%s", header, m.viewport.View(), footer)
}

// noteListModel is an interactive list for browsing notes
type noteListModel struct {
	list         list.Model
	viewport     viewport.Model
	showPreview  bool
	notesDir     string
	sortMode     string // "name", "date", "size"
	allNotes     []noteInfo
	quitting     bool
	selected     *noteInfo
	action       string // "", "create", "rename", "delete"
	width        int
	height       int
}

func newNoteListModel(notes []noteInfo, sortMode string) noteListModel {
	// Convert notes to list items
	items := make([]list.Item, len(notes))
	for i, note := range notes {
		items[i] = note
	}

	// Create list with custom delegate for styling
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(theme.Primary.GetForeground()).
		BorderForeground(theme.Accent.GetForeground())
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(theme.Secondary.GetForeground())

	l := list.New(items, delegate, 0, 0)
	l.Title = "Notes"
	l.Styles.Title = theme.Header
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)

	// Add custom keybindings help
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "open")),
			key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new")),
			key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "rename")),
			key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "sort")),
			key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "preview")),
		}
	}

	vp := viewport.New(0, 0)

	notesDir, _ := getNotesDir()

	return noteListModel{
		list:        l,
		viewport:    vp,
		showPreview: true, // Preview visible by default
		notesDir:    notesDir,
		sortMode:    sortMode,
		allNotes:    notes,
		quitting:    false,
		selected:    nil,
		action:      "",
		width:       0,
		height:      0,
	}
}

func (m noteListModel) Init() tea.Cmd {
	return nil
}

func (m noteListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "ctrl+c":
			m.action = "quit"
			m.quitting = true
			return m, tea.Quit

		case "enter":
			// Open selected note
			if item, ok := m.list.SelectedItem().(noteInfo); ok {
				m.selected = &item
				m.action = "open"
				m.quitting = true
				return m, tea.Quit
			}

		case "n":
			// Create new note
			m.action = "create"
			m.quitting = true
			return m, tea.Quit

		case "e":
			// Rename selected note
			if item, ok := m.list.SelectedItem().(noteInfo); ok {
				m.selected = &item
				m.action = "rename"
				m.quitting = true
				return m, tea.Quit
			}

		case "d":
			// Delete selected note
			if item, ok := m.list.SelectedItem().(noteInfo); ok {
				m.selected = &item
				m.action = "delete"
				m.quitting = true
				return m, tea.Quit
			}

		case "s":
			// Cycle sort mode
			switch m.sortMode {
			case "name":
				m.sortMode = "date"
			case "date":
				m.sortMode = "size"
			case "size":
				m.sortMode = "name"
			}
			// Re-sort and update list
			m.allNotes = sortNotes(m.allNotes, m.sortMode)
			items := make([]list.Item, len(m.allNotes))
			for i, note := range m.allNotes {
				items[i] = note
			}
			m.list.SetItems(items)
			m.list.Title = fmt.Sprintf("Notes (sorted by: %s)", m.sortMode)
			return m, nil

		case "p":
			// Toggle preview
			m.showPreview = !m.showPreview
			// Force resize to recalculate layout
			if m.width > 0 && m.height > 0 {
				return m.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		h, v := lipgloss.NewStyle().GetFrameSize()

		if m.showPreview {
			// Split view: list on left, preview on right
			listWidth := msg.Width / 2
			previewWidth := msg.Width - listWidth - 2 // -2 for border

			m.list.SetSize(listWidth-h, msg.Height-v)
			m.viewport.Width = previewWidth
			m.viewport.Height = msg.Height - v - 3 // -3 for header

			// Update preview content
			if item, ok := m.list.SelectedItem().(noteInfo); ok {
				content, err := os.ReadFile(filepath.Join(m.notesDir, item.name))
				if err == nil {
					m.viewport.SetContent(string(content))
				} else {
					m.viewport.SetContent(theme.Error.Render("Error reading file"))
				}
			}
		} else {
			m.list.SetSize(msg.Width-h, msg.Height-v)
		}
	}

	var cmd tea.Cmd

	// Update list and possibly viewport
	m.list, cmd = m.list.Update(msg)

	// If preview is shown and selection changed, update preview
	if m.showPreview {
		if item, ok := m.list.SelectedItem().(noteInfo); ok {
			content, err := os.ReadFile(filepath.Join(m.notesDir, item.name))
			if err == nil {
				m.viewport.SetContent(string(content))
			}
		}
	}

	return m, cmd
}

func (m noteListModel) View() string {
	if m.quitting {
		return ""
	}

	if m.showPreview {
		// Split view: list on left, preview on right
		previewHeader := theme.Header.Render(" Preview ")
		previewContent := m.viewport.View()
		previewPanel := lipgloss.JoinVertical(lipgloss.Left, previewHeader, previewContent)

		previewStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.Accent.GetForeground()).
			Padding(0, 1)

		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.list.View(),
			previewStyle.Render(previewPanel),
		)
	}

	return m.list.View()
}

// menuModel is the main menu
type menuModel struct {
	choices  []string
	cursor   int
	selected string
	quitting bool
	width    int
	height   int
}

func newMenuModel() menuModel {
	return menuModel{
		choices: []string{
			"Browse Notes",
			"Search Notes",
			"Create New Note",
			"Change Theme",
			"Help",
			"Quit",
		},
		cursor: 0,
		width:  0,
		height: 0,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			m.selected = "quit"
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = m.choices[m.cursor]
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	if m.quitting {
		return ""
	}

	// Header
	header := theme.Header.Render(" ks - Keep Simple Notes ")

	// Build menu items
	var menuItems strings.Builder
	menuItems.WriteString("\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "›"
			menuItems.WriteString(theme.Selected.Render(cursor + " " + choice))
		} else {
			menuItems.WriteString(theme.Muted.Render(cursor + " " + choice))
		}
		menuItems.WriteString("\n")
	}

	// Footer
	footer := "\n" + theme.Muted.Render("↑/↓: navigate • enter: select • q: quit")

	// Center the content vertically
	content := header + menuItems.String() + footer

	// Calculate vertical centering
	contentHeight := strings.Count(content, "\n") + 1
	topPadding := 0
	if m.height > contentHeight {
		topPadding = (m.height - contentHeight) / 2
	}

	// Apply vertical centering
	if topPadding > 0 {
		content = strings.Repeat("\n", topPadding) + content
	}

	// Center horizontally with full width
	style := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center)

	return style.Render(content)
}

// themeSelectModel handles theme selection
type themeSelectModel struct {
	themeNames []string
	cursor     int
	selected   string
	quitting   bool
	width      int
	height     int
}

func newThemeSelectModel() themeSelectModel {
	themeNames := []string{"Purple (Default)", "Ocean", "Forest", "Sunset"}

	// Find current theme's index
	cursor := 0
	for i, name := range themeNames {
		if name == currentThemeName {
			cursor = i
			break
		}
	}

	return themeSelectModel{
		themeNames: themeNames,
		cursor:     cursor,
		quitting:   false,
		width:      0,
		height:     0,
	}
}

func (m themeSelectModel) Init() tea.Cmd {
	return nil
}

func (m themeSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.themeNames)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = m.themeNames[m.cursor]
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m themeSelectModel) View() string {
	if m.quitting {
		return ""
	}

	// Header
	header := theme.Header.Render(" Choose Theme ")

	// Description
	description := "\n" + theme.Secondary.Render("Select a theme to change the application's color scheme") + "\n"

	// Build theme list
	var themeList strings.Builder
	themeList.WriteString("\n")
	for i, themeName := range m.themeNames {
		cursor := " "
		currentMarker := ""

		if themeName == currentThemeName {
			currentMarker = " " + theme.Success.Render("(current)")
		}

		if m.cursor == i {
			cursor = "›"
			themeList.WriteString(theme.Selected.Render(cursor+" "+themeName) + currentMarker)
		} else {
			themeList.WriteString(theme.Muted.Render(cursor+" "+themeName) + currentMarker)
		}
		themeList.WriteString("\n")
	}

	// Footer
	footer := "\n" + theme.Muted.Render("↑/↓: navigate • enter: select • q: cancel")

	// Combine all content
	content := header + description + themeList.String() + footer

	// Calculate vertical centering
	contentHeight := strings.Count(content, "\n") + 1
	topPadding := 0
	if m.height > contentHeight {
		topPadding = (m.height - contentHeight) / 2
	}

	// Apply vertical centering
	if topPadding > 0 {
		content = strings.Repeat("\n", topPadding) + content
	}

	// Center horizontally with full width
	style := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center)

	return style.Render(content)
}

// runThemeSelector launches the theme selection UI
func runThemeSelector() {
	m := newThemeSelectModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	result, err := p.Run()

	if err != nil {
		fmt.Println(theme.Error.Render("✗ Error: " + err.Error()))
		return
	}

	themeSelect := result.(themeSelectModel)

	if themeSelect.selected != "" {
		// Apply the selected theme
		if themeFn, ok := themes[themeSelect.selected]; ok {
			theme = themeFn()
			currentThemeName = themeSelect.selected

			// Update legacy style aliases
			selectedStyle = theme.Selected
			unselectedStyle = theme.Unselected
			boxStyle = theme.Border

			fmt.Println(theme.Success.Render("✓ Theme changed to: " + currentThemeName))
			fmt.Println(theme.Muted.Render("\nPress Enter to continue..."))
			fmt.Scanln()
		}
	}
}

// writeNote writes a note to a file
func writeNote(filename, note string) {
	// Validate filename first
	if err := validateFilename(filename); err != nil {
		fmt.Printf("Invalid filename: %v\n", err)
		os.Exit(1)
	}

	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Build the full file path
	filePath := filepath.Join(notesDir, filename)

	// Write the note to the file
	err = os.WriteFile(filePath, []byte(note), 0644)
	if err != nil {
		fmt.Println(theme.Error.Render("✗ Error writing file: " + err.Error()))
		os.Exit(1)
	}

	fmt.Println(theme.Success.Render("✓ Successfully wrote note to " + filePath))
}

// appendNote appends content to an existing note (or creates it if it doesn't exist)
func appendNote(filename, note string) {
	// Validate filename first
	if err := validateFilename(filename); err != nil {
		fmt.Printf("Invalid filename: %v\n", err)
		os.Exit(1)
	}

	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Build the full file path
	filePath := filepath.Join(notesDir, filename)

	// Check if file exists and get its size
	fileInfo, err := os.Stat(filePath)
	needsNewline := false

	// If file doesn't exist, ask for confirmation to create it
	if os.IsNotExist(err) {
		if !confirm(fmt.Sprintf("File '%s' does not exist. Create it?", filename)) {
			fmt.Println("Append cancelled.")
			return
		}
		// User confirmed, proceed with creation (needsNewline stays false for new files)
	} else if err == nil && fileInfo.Size() > 0 {
		// File exists and has content - check if it ends with newline
		// Read the last byte to check if it's a newline
		file, err := os.Open(filePath)
		if err != nil {
			// If we can't open to check, assume we need a newline to be safe
			needsNewline = true
		} else {
			defer file.Close()
			// Seek to the last byte
			_, err = file.Seek(-1, io.SeekEnd)
			if err != nil {
				// If we can't seek, assume we need a newline
				needsNewline = true
			} else {
				lastByte := make([]byte, 1)
				_, err = file.Read(lastByte)
				if err != nil {
					// If we can't read, assume we need a newline
					needsNewline = true
				} else if lastByte[0] != '\n' {
					// Last byte is not a newline, we need to add one
					needsNewline = true
				}
			}
		}
	}

	// Open file with append mode, create if doesn't exist, write-only
	// O_APPEND: Append to end of file
	// O_CREATE: Create file if it doesn't exist
	// O_WRONLY: Write-only access
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	// defer ensures file is closed when function exits (even if there's an error)
	defer file.Close()

	// If the file doesn't end with newline, add one before appending
	if needsNewline {
		_, err = file.WriteString("\n")
		if err != nil {
			fmt.Printf("Error adding newline: %v\n", err)
			os.Exit(1)
		}
	}

	// Write the note to the file
	_, err = file.WriteString(note)
	if err != nil {
		fmt.Println(theme.Error.Render("✗ Error appending to file: " + err.Error()))
		os.Exit(1)
	}

	fmt.Println(theme.Success.Render("✓ Successfully appended to " + filePath))
}

// noteInfo holds information about a note file for sorting
type noteInfo struct {
	name    string
	modTime time.Time
	size    int64
}

// Implement list.Item interface for noteInfo
func (n noteInfo) FilterValue() string { return n.name }
func (n noteInfo) Title() string       { return n.name }
func (n noteInfo) Description() string {
	timeStr := n.modTime.Format("2006-01-02 15:04")
	sizeStr := formatSize(n.size)
	return fmt.Sprintf("%s • %s", sizeStr, timeStr)
}

// searchResult holds information about a search match
type searchResult struct {
	note          noteInfo
	matchLocation string // "filename", "content", or "filename and content"
}

// Implement list.Item interface for searchResult
func (s searchResult) FilterValue() string { return s.note.name }
func (s searchResult) Title() string       { return s.note.name }
func (s searchResult) Description() string {
	return fmt.Sprintf("Match in: %s", s.matchLocation)
}

// handleListAction processes actions returned from the note list
func handleListAction(m noteListModel) {
	switch m.action {
	case "open":
		if m.selected != nil {
			readNote(m.selected.name)
		}
	case "create":
		runInteractiveCreate()
	case "rename":
		if m.selected != nil {
			runInteractiveRename(m.selected.name)
		}
	case "delete":
		if m.selected != nil {
			deleteNote(m.selected.name, false)
		}
	case "quit":
		// Do nothing, return to REPL
	}
}

// runInteractiveRename prompts for a new filename and renames the note
func runInteractiveRename(oldName string) {
	ti := textinput.New()
	ti.Placeholder = oldName
	ti.SetValue(oldName)
	ti.Focus()

	fmt.Print(theme.Primary.Render("New filename: "))

	var newName string
	fmt.Scanln(&newName)

	if newName != "" && newName != oldName {
		if err := validateFilename(newName); err != nil {
			fmt.Println(theme.Error.Render("✗ Invalid filename: " + err.Error()))
			return
		}

		notesDir, _ := getNotesDir()
		oldPath := filepath.Join(notesDir, oldName)
		newPath := filepath.Join(notesDir, newName)

		if err := os.Rename(oldPath, newPath); err != nil {
			fmt.Println(theme.Error.Render("✗ Error renaming: " + err.Error()))
		} else {
			fmt.Println(theme.Success.Render("✓ Renamed to " + newName))
		}
	}
}

// sortNotes sorts a slice of noteInfo by the specified mode
func sortNotes(notes []noteInfo, sortBy string) []noteInfo {
	sorted := make([]noteInfo, len(notes))
	copy(sorted, notes)

	switch sortBy {
	case "date":
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].modTime.After(sorted[j].modTime)
		})
	case "size":
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].size > sorted[j].size
		})
	default: // "name"
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].name < sorted[j].name
		})
	}

	return sorted
}

// listNotes lists all notes in the notes directory with optional sorting
func listNotes(sortBy string, interactive bool) {
	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Read all entries in the notes directory
	entries, err := os.ReadDir(notesDir)
	if err != nil {
		fmt.Printf("Error reading notes directory: %v\n", err)
		os.Exit(1)
	}

	// Collect file information into a slice
	var notes []noteInfo
	for _, entry := range entries {
		// Skip directories, only process files
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				// If we can't get info, skip this file
				continue
			}

			notes = append(notes, noteInfo{
				name:    entry.Name(),
				modTime: info.ModTime(),
				size:    info.Size(),
			})
		}
	}

	// Check if there are any notes
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	// Sort the notes based on the sortBy parameter
	switch sortBy {
	case "date":
		// Sort by modification time (newest first)
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].modTime.After(notes[j].modTime)
		})
	case "size":
		// Sort by size (largest first)
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].size > notes[j].size
		})
	default:
		// Sort by name (alphabetical)
		sort.Slice(notes, func(i, j int) bool {
			return notes[i].name < notes[j].name
		})
	}

	// If interactive mode, launch TUI list
	if interactive && isTTY() {
		m := newNoteListModel(notes, sortBy)
		p := tea.NewProgram(m, tea.WithAltScreen())

		finalModel, err := p.Run()
		if err != nil {
			fmt.Println(theme.Error.Render("✗ Error running list: " + err.Error()))
			os.Exit(1)
		}

		// Handle actions from the list
		if final, ok := finalModel.(noteListModel); ok {
			handleListAction(final)
		}
		return
	}

	// Non-interactive mode: simple list display
	fmt.Println(theme.Header.Render("Notes:"))
	// Display the sorted notes
	for _, note := range notes {
		timeStr := note.modTime.Format("2006-01-02 15:04")
		sizeStr := formatSize(note.size)
		nameStyled := theme.Primary.Render(note.name)
		metaStyled := theme.Secondary.Render(fmt.Sprintf("%8s  (modified: %s)", sizeStr, timeStr))
		fmt.Printf("  • %s %s\n", nameStyled, metaStyled)
	}
}

// formatSize formats file size in human-readable format
func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// readNote reads and displays a note
func readNote(filename string) {
	// Validate filename first
	if err := validateFilename(filename); err != nil {
		fmt.Printf("Invalid filename: %v\n", err)
		os.Exit(1)
	}

	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Build the full file path
	filePath := filepath.Join(notesDir, filename)

	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Note '%s' not found.\n", filename)
		} else {
			fmt.Printf("Error reading file: %v\n", err)
		}
		os.Exit(1)
	}

	// Launch interactive viewport for reading the note
	// Check if we have a TTY - if not, fall back to simple print
	if !isTTY() {
		// Fallback for non-TTY environments (pipes, redirects)
		header := theme.Header.Render(" " + filename + " ")
		fmt.Println(header)
		fmt.Println(string(content))
		return
	}

	m := newNoteViewerModel(filename, string(content))
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(theme.Error.Render("✗ Error running viewer: " + err.Error()))
		os.Exit(1)
	}
}

// deleteNote deletes a note
func deleteNote(filename string, force bool) {
	// Validate filename first
	if err := validateFilename(filename); err != nil {
		fmt.Printf("Invalid filename: %v\n", err)
		os.Exit(1)
	}

	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Build the full file path
	filePath := filepath.Join(notesDir, filename)

	// Check if file exists before asking for confirmation
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Note '%s' not found.\n", filename)
		os.Exit(1)
	}

	// Ask for confirmation unless --force is used
	if !force {
		if !confirm(fmt.Sprintf("Delete '%s'?", filename)) {
			fmt.Println("Deletion cancelled.")
			return
		}
	}

	// Delete the file
	err = os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(theme.Error.Render("✗ Note '" + filename + "' not found"))
		} else {
			fmt.Println(theme.Error.Render("✗ Error deleting file: " + err.Error()))
		}
		os.Exit(1)
	}

	fmt.Println(theme.Success.Render("✓ Successfully deleted note: " + filename))
}

// searchNotes searches for a keyword in all notes (filenames and content)
func searchNotes(keyword string, interactive bool) {
	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Read all entries in the notes directory
	entries, err := os.ReadDir(notesDir)
	if err != nil {
		fmt.Printf("Error reading notes directory: %v\n", err)
		os.Exit(1)
	}

	// Convert keyword to lowercase for case-insensitive search
	keywordLower := strings.ToLower(keyword)
	var results []searchResult

	// Search through each file
	for _, entry := range entries {
		// Skip directories, only process files
		if !entry.IsDir() {
			filePath := filepath.Join(notesDir, entry.Name())
			filenameLower := strings.ToLower(entry.Name())

			// Check if filename matches
			filenameMatch := strings.Contains(filenameLower, keywordLower)

			// Read the file content
			content, err := os.ReadFile(filePath)
			contentMatch := false

			// Check content if we can read the file
			if err == nil {
				contentLower := strings.ToLower(string(content))
				contentMatch = strings.Contains(contentLower, keywordLower)
			}

			// If either filename or content matches, add to results
			if filenameMatch || contentMatch {
				info, _ := entry.Info()
				matchLocation := "content"
				if filenameMatch && contentMatch {
					matchLocation = "filename and content"
				} else if filenameMatch {
					matchLocation = "filename"
				}

				results = append(results, searchResult{
					note: noteInfo{
						name:    entry.Name(),
						modTime: info.ModTime(),
						size:    info.Size(),
					},
					matchLocation: matchLocation,
				})
			}
		}
	}

	// Check if any results found
	if len(results) == 0 {
		fmt.Println(theme.Primary.Render("Searching for: ") + theme.Accent.Render(keyword))
		fmt.Println()
		fmt.Println(theme.Warning.Render("No matches found."))
		return
	}

	// If interactive mode and TTY available, show interactive list
	if interactive && isTTY() {
		// Convert search results to notes for the list model
		notes := make([]noteInfo, len(results))
		for i, result := range results {
			notes[i] = result.note
		}

		m := newNoteListModel(notes, "name")
		m.list.Title = fmt.Sprintf("Search Results for: %s", keyword)

		p := tea.NewProgram(m, tea.WithAltScreen())
		finalModel, err := p.Run()
		if err != nil {
			fmt.Println(theme.Error.Render("✗ Error running search: " + err.Error()))
			os.Exit(1)
		}

		// Handle actions from the list
		if final, ok := finalModel.(noteListModel); ok {
			handleListAction(final)
		}
		return
	}

	// Non-interactive mode: simple list display
	fmt.Println(theme.Primary.Render("Searching for: ") + theme.Accent.Render(keyword))
	fmt.Println()

	for _, result := range results {
		matchLocation := theme.Accent.Render(result.matchLocation)
		nameStyled := theme.Primary.Render(result.note.name)
		fmt.Printf("  • %s %s\n", nameStyled, theme.Secondary.Render("(match in: ")+matchLocation+theme.Secondary.Render(")"))
	}

	fmt.Println()
	fmt.Println(theme.Success.Render(fmt.Sprintf("✓ Found %d match(es)", len(results))))
}
