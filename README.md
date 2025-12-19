##🛡️ Security Toolkit App (Go)
##✨ Short Description

The Security Toolkit App is a menu-driven command-line application built in Go (Golang) that demonstrates core cybersecurity concepts. It combines cryptography and file integrity verification into a single interactive tool.

The application allows users to:

Encrypt and decrypt messages using the Rail Fence Cipher

Generate and verify SHA-256 hashes to detect file modification, corruption, or tampering

This project is designed for educational and academic purposes, helping students understand how encryption and integrity mechanisms work in practice.

##📦 Dependencies or Libraries Used

This project uses only the Go Standard Library, ensuring simplicity, portability, and no external dependencies.

Library	Purpose
fmt	Displays menus, messages, and formatted output
bufio	Handles user input from the command line
os	File handling and system interaction
strings	String manipulation and input processing
strconv	Converts user input (e.g., cipher key)
crypto/sha256	Generates secure SHA-256 hashes
encoding/hex	Converts hash bytes to readable hexadecimal
io	Streams file data efficiently for hashing
##🛠️ Installation / Setup Instructions
Prerequisites

Go (Golang)
Make sure Go is installed:

go version


Operating System
Works on Windows, Linux, and macOS.

Setup

Clone or Download the Project

git clone <repository-url>


Navigate to the Project Directory

cd security-toolkit-app


Run the Application

go run *.go

##🚀 Usage Examples

When the program starts, the following menu is displayed:

+----------------------------------+
|        Security Toolkit App       |
+----------------------------------+
| 1. Cipher Module                 |
| 2. File Integrity Checker        |
| 0. Exit                          |
+----------------------------------+
Choose an option:

##🔐 Cipher Module
Encrypt a Message (Rail Fence Cipher)

Steps:

Choose option 1

Select Encrypt Message

Enter the message

Enter the number of rails (key)

Example:

Enter message: HELLO WORLD
Enter key (number of rails): 3
Encrypted Message: HOREL OLLWD

Decrypt a Message

Example:

Enter message: HOREL OLLWD
Enter key (number of rails): 3
Decrypted Message: HELLO WORLD

##📁 File Integrity Checker
1. Generate SHA-256 Hash

Steps:

Choose option 2

Select Generate SHA-256 Hash

Enter file path

Example:

Enter file path: report.txt
SHA-256 Hash:
a7c3f1b2e9...

2. Verify File Integrity

Example (File Intact):

Enter file path: report.txt
Enter known hash: a7c3f1b2e9...
✅ File is intact


Example (File Modified):

❌ File has been modified

##🔄 Program Flow (High-Level)
Start Program
   ↓
Display Main Menu
   ↓
User Selects Option
   ↓
Cipher Module → Encrypt / Decrypt
   ↓
File Integrity Checker → Hash / Verify
   ↓
Return to Menu
   ↓
Exit

##🎯 Learning Outcomes

Understand classical encryption (Rail Fence Cipher)

Learn how file integrity is verified using cryptographic hashes

Practice building menu-driven CLI applications in Go

Gain hands-on experience with security-focused programming

##📌 Conclusion

The Security Toolkit App provides a simple yet effective way to explore foundational security concepts through hands-on interaction. By combining cryptography and file integrity verification, this tool bridges theory and practice, making it ideal for cybersecurity students and beginners in Go programming.
