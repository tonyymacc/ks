package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if we have at least one argument (the subcommand)
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Get the subcommand (first argument)
	subcommand := os.Args[1]

	// Execute the appropriate function based on subcommand
	switch subcommand {
	case "write":
		if len(os.Args) != 4 {
			fmt.Println("Usage: simplecli write <filename> <note>")
			os.Exit(1)
		}
		writeNote(os.Args[2], os.Args[3])
	case "list":
		if len(os.Args) != 2 {
			fmt.Println("Usage: simplecli list")
			os.Exit(1)
		}
		listNotes()
	case "read":
		if len(os.Args) != 3 {
			fmt.Println("Usage: simplecli read <filename>")
			os.Exit(1)
		}
		readNote(os.Args[2])
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: simplecli delete <filename>")
			os.Exit(1)
		}
		deleteNote(os.Args[2])
	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

// printUsage displays the help message
func printUsage() {
	fmt.Println("Usage: simplecli <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  write <filename> <note>  - Write a note to a file")
	fmt.Println("  list                     - List all notes")
	fmt.Println("  read <filename>          - Read a note")
	fmt.Println("  delete <filename>        - Delete a note")
}

// getNotesDir returns the path to the notes directory
func getNotesDir() (string, error) {
	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Build the notes directory path: ~/.local/share/simplecli
	notesDir := filepath.Join(home, ".local", "share", "simplecli")

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
