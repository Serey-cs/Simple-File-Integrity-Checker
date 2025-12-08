package main

import (
    "fmt"
    "os"
)

func main() {
    // Check arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage:")
        fmt.Println("  Generate hash: go run *.go <file_path>")
        fmt.Println("  Verify hash:   go run *.go <file_path> <known_hash>")
        return
    }

    filePath := os.Args[1]

    // If a known hash is provided → verification mode
    if len(os.Args) >= 3 {
        knownHash := os.Args[2]
        match, err := verifyFile(filePath, knownHash)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        if match {
            fmt.Println("✅ File is intact")
        } else {
            fmt.Println("❌ File has been modified or corrupted")
        }
        return
    }

    // No known hash → generate SHA-256
    hash, err := computeSHA256(filePath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("SHA-256 hash of %s:\n%s\n", filePath, hash)
}
