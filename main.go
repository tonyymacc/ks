package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	// Define flag variables (will be set to true if flag is used)
	var writeFlag, listFlag, readFlag, deleteFlag, appendFlag, searchFlag bool

	// Register short flags (-w, -l, -r, -d, -a, -s)
	flag.BoolVar(&writeFlag, "w", false, "Write a note")
	flag.BoolVar(&listFlag, "l", false, "List all notes")
	flag.BoolVar(&readFlag, "r", false, "Read a note")
	flag.BoolVar(&deleteFlag, "d", false, "Delete a note")
	flag.BoolVar(&appendFlag, "a", false, "Append to a note")
	flag.BoolVar(&searchFlag, "s", false, "Search notes")

	// Register long flags (--write, --list, --read, --delete, --append, --search)
	flag.BoolVar(&writeFlag, "write", false, "Write a note")
	flag.BoolVar(&listFlag, "list", false, "List all notes")
	flag.BoolVar(&readFlag, "read", false, "Read a note")
	flag.BoolVar(&deleteFlag, "delete", false, "Delete a note")
	flag.BoolVar(&appendFlag, "append", false, "Append to a note")
	flag.BoolVar(&searchFlag, "search", false, "Search notes")

	// Sorting flags (for list command)
	var sortBy string
	flag.StringVar(&sortBy, "sort", "name", "Sort order for list: name, date, size")

	// Custom usage message
	flag.Usage = printUsage

	// Parse the flags
	flag.Parse()

	// Get remaining arguments after flags
	args := flag.Args()

	// If no flags provided, show usage (later will launch TUI)
	if flag.NFlag() == 0 {
		printUsage()
		os.Exit(1)
	}

	// Check that only one flag is used at a time
	flagCount := 0
	if writeFlag {
		flagCount++
	}
	if listFlag {
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
	if searchFlag {
		flagCount++
	}

	if flagCount > 1 {
		fmt.Println("Error: Only one command flag can be used at a time")
		printUsage()
		os.Exit(1)
	}

	// Execute the appropriate command based on flag
	if writeFlag {
		if len(args) != 2 {
			fmt.Println("Usage: ks -w <filename> <note>")
			fmt.Println("   or: ks --write <filename> <note>")
			os.Exit(1)
		}
		writeNote(args[0], args[1])
	} else if listFlag {
		if len(args) != 0 {
			fmt.Println("Usage: ks -l [--sort name|date|size]")
			fmt.Println("   or: ks --list [--sort name|date|size]")
			os.Exit(1)
		}
		listNotes(sortBy)
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
		deleteNote(args[0])
	} else if appendFlag {
		if len(args) != 2 {
			fmt.Println("Usage: ks -a <filename> <note>")
			fmt.Println("   or: ks --append <filename> <note>")
			os.Exit(1)
		}
		appendNote(args[0], args[1])
	} else if searchFlag {
		if len(args) != 1 {
			fmt.Println("Usage: ks -s <keyword>")
			fmt.Println("   or: ks --search <keyword>")
			os.Exit(1)
		}
		searchNotes(args[0])
	}
}

// printUsage displays the help message
func printUsage() {
	fmt.Println("ks - Keep It Simple Stupid")
	fmt.Println("\nUsage:")
	fmt.Println("  ks [flags] [arguments]")
	fmt.Println("\nFlags:")
	fmt.Println("  -w, --write <filename> <note>    Write a note to a file")
	fmt.Println("  -a, --append <filename> <note>   Append to an existing note")
	fmt.Println("  -l, --list [--sort order]        List all notes")
	fmt.Println("  -r, --read <filename>            Read a note")
	fmt.Println("  -d, --delete <filename>          Delete a note")
	fmt.Println("  -s, --search <keyword>           Search notes for keyword")
	fmt.Println("\nList Options:")
	fmt.Println("  --sort name     Sort by filename (default)")
	fmt.Println("  --sort date     Sort by modification time (newest first)")
	fmt.Println("  --sort size     Sort by file size (largest first)")
	fmt.Println("\nExamples:")
	fmt.Println("  ks -w note.txt \"My note content\"")
	fmt.Println("  ks -a note.txt \"\\nMore content\"")
	fmt.Println("  ks -l")
	fmt.Println("  ks -l --sort date")
	fmt.Println("  ks -r note.txt")
	fmt.Println("  ks -s golang")
	fmt.Println("  ks -d note.txt")
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

// writeNote writes a note to a file
func writeNote(filename, note string) {
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
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully wrote note to %s\n", filePath)
}

// appendNote appends content to an existing note (or creates it if it doesn't exist)
func appendNote(filename, note string) {
	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Build the full file path
	filePath := filepath.Join(notesDir, filename)

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

	// Write the note to the file
	_, err = file.WriteString(note)
	if err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully appended to %s\n", filePath)
}

// noteInfo holds information about a note file for sorting
type noteInfo struct {
	name    string
	modTime time.Time
	size    int64
}

// listNotes lists all notes in the notes directory with optional sorting
func listNotes(sortBy string) {
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

	fmt.Println("Notes:")
	// Display the sorted notes
	for _, note := range notes {
		timeStr := note.modTime.Format("2006-01-02 15:04")
		sizeStr := formatSize(note.size)
		fmt.Printf("  - %-30s %8s  (modified: %s)\n", note.name, sizeStr, timeStr)
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

	// Display the filename and content
	fmt.Printf("=== %s ===\n", filename)
	fmt.Println(string(content))
}

// deleteNote deletes a note
func deleteNote(filename string) {
	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error getting notes directory: %v\n", err)
		os.Exit(1)
	}

	// Build the full file path
	filePath := filepath.Join(notesDir, filename)

	// Delete the file
	err = os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Note '%s' not found.\n", filename)
		} else {
			fmt.Printf("Error deleting file: %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("Successfully deleted note: %s\n", filename)
}

// searchNotes searches for a keyword in all notes (filenames and content)
func searchNotes(keyword string) {
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
	matchCount := 0

	fmt.Printf("Searching for: %s\n\n", keyword)

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

			// If either filename or content matches, show the file
			if filenameMatch || contentMatch {
				matchCount++

				// Determine match location
				var matchLocation string
				if filenameMatch && contentMatch {
					matchLocation = "filename and content"
				} else if filenameMatch {
					matchLocation = "filename"
				} else {
					matchLocation = "content"
				}

				fmt.Printf("  - %-30s (match in: %s)\n", entry.Name(), matchLocation)
			}
		}
	}

	if matchCount == 0 {
		fmt.Println("No matches found.")
	} else {
		fmt.Printf("\nFound %d match(es).\n", matchCount)
	}
}
