package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define flag variables (will be set to true if flag is used)
	var writeFlag, listFlag, readFlag, deleteFlag, appendFlag bool

	// Register short flags (-w, -l, -r, -d, -a)
	flag.BoolVar(&writeFlag, "w", false, "Write a note")
	flag.BoolVar(&listFlag, "l", false, "List all notes")
	flag.BoolVar(&readFlag, "r", false, "Read a note")
	flag.BoolVar(&deleteFlag, "d", false, "Delete a note")
	flag.BoolVar(&appendFlag, "a", false, "Append to a note")

	// Register long flags (--write, --list, --read, --delete, --append)
	flag.BoolVar(&writeFlag, "write", false, "Write a note")
	flag.BoolVar(&listFlag, "list", false, "List all notes")
	flag.BoolVar(&readFlag, "read", false, "Read a note")
	flag.BoolVar(&deleteFlag, "delete", false, "Delete a note")
	flag.BoolVar(&appendFlag, "append", false, "Append to a note")

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
			fmt.Println("Usage: ks -l")
			fmt.Println("   or: ks --list")
			os.Exit(1)
		}
		listNotes()
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
	fmt.Println("  -l, --list                       List all notes")
	fmt.Println("  -r, --read <filename>            Read a note")
	fmt.Println("  -d, --delete <filename>          Delete a note")
	fmt.Println("\nExamples:")
	fmt.Println("  ks -w note.txt \"My note content\"")
	fmt.Println("  ks -a note.txt \"\\nMore content\"")
	fmt.Println("  ks -l")
	fmt.Println("  ks -r note.txt")
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

// listNotes lists all notes in the notes directory
func listNotes() {
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

	// Check if there are any notes
	if len(entries) == 0 {
		fmt.Println("No notes found.")
		return
	}

	fmt.Println("Notes:")
	// Loop through each entry
	for _, entry := range entries {
		// Skip directories, only show files
		if !entry.IsDir() {
			fmt.Printf("  - %s\n", entry.Name())
		}
	}
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
