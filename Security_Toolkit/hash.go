package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io"
    "os"
)

// computeSHA256 takes a file path and returns its SHA-256 hash as a string
func computeSHA256(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hasher := sha256.New()

    if _, err := io.Copy(hasher, file); err != nil {
        return "", err
    }

    hashBytes := hasher.Sum(nil)

    return hex.EncodeToString(hashBytes), nil
}
