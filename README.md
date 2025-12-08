# 🐍 Simple File Integrity Checker (FIC)

## ✨ Short Description

This command-line utility, built in **Go (Golang)**, provides a simple, robust solution for ensuring **data integrity**. It generates cryptographic hashes (**SHA-256**) for any given file. By comparing a newly calculated hash against a known, trusted hash, users can quickly verify that a file has not been **modified, corrupted, or tampered with** since its initial creation or transfer. 

---

## 📦 Dependencies or Libraries Used

This project uses only the **Go Standard Library**, ensuring minimal setup and high performance.

| Library | Purpose |
| :--- | :--- |
| `crypto/sha256` | Provides the necessary cryptographic functions to calculate the SHA-256 hash. |
| `io` | Used for efficient **streaming** of file content into the hash function, avoiding excessive memory use for large files. |
| `os` | Used for file operations (opening files) and command-line argument handling. |
| `fmt` | Used for formatted output (e.g., printing the hex-encoded hash). |

---

## 🛠️ Installation/Setup Instructions

### Prerequisites

1.  **Go:** Ensure you have Go installed on your system. You can verify this by running `go version` in your terminal.
2.  **Version Control (Bonus):** It's recommended to initialize a Git repository for clean version tracking.

### Setup

1.  **Save the Code:** Save the entire program code into a single folder named `fic`.
2.  **Navigate:** Open your terminal and navigate to the directory where you saved `fic`.

---

## 🚀 Usage Examples

The tool supports two primary functions: `generate` (to get the hash) and `verify` (to check the hash).

### 1. Generating a Hash

To generate and print the SHA-256 hash for a specific file:

### Terminal

go run main.go hash.go utils.go verify.go [path/to/your/file.txt]

---

### Example Output:

SHA-256 hash of Hello.txt: 5f4e03d4b6c3f2a1e8d7b9c0a2f1e6d5b4c3a2b1e0d9c8b7a6f5e4d3c2b1a0f9

---

### 2. Verifying a Hash

To verify a file against a known hash value (e.g., one provided by a software vendor):

go run main.go hash.go utils.go verify.go [path/to/your/file.txt] [known_hash_value]

---

### Example of Success:

go run main.go hash.go utils.go verify.go Hello.txt 1a2b3c4d5e6f...
# Output: ✅ File is intact

---

### Example of Failure (File Tampered/Corrupted):

go run main.go hash.go utils.go verify.go Hello.txt000000000000...
# Output: ❌ File has been modified or corrupted