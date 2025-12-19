package main

import "fmt"

// verifyFile checks if the SHA-256 hash of a file matches the given hash
func verifyFile(path string, knownHash string) (bool, error) {
    hash, err := computeSHA256(path)
    if err != nil {
        return false, err
    }

    return hash == knownHash, nil
}

// optional helper to print results in a friendly way
func printVerificationResult(match bool, path string) {
    if match {
        fmt.Printf("✅ File is good: %s\n", path)
    } else {
        fmt.Printf("❌ File has changed: %s\n", path)
    }
}
