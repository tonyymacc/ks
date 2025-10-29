package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if we have exactly 2 arguments (plus program name = 3 total)
	if len(os.Args) != 3 {
		fmt.Println("Usage: simplecli <filename> <note>")
		os.Exit(1)
	}

	// Get the arguments
	filename := os.Args[1]
	note := os.Args[2]

	// Get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	// Build the notes directory path: ~/.local/share/simplecli
	notesDir := filepath.Join(home, ".local", "share", "simplecli")

	// Create the directory if it doesn't exist (including parent directories)
	err = os.MkdirAll(notesDir, 0755)
	if err != nil {
		fmt.Printf("Error creating notes directory: %v\n", err)
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
