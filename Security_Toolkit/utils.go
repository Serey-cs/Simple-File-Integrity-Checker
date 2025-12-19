package main

import "fmt"

// prints usage instructions
func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  go run *.go generate <file_path>        → get SHA-256 of a file")
    fmt.Println("  go run *.go verify <file_path> <hash>   → check file against a known hash")
}

// prints an error message (simple logging)
func logError(err error) {
    fmt.Println("Error:", err)
}
