<<<<<<< HEAD
# 🛡️ Security Toolkit App (Go)

## ✨ Overview

The **Security Toolkit App** is a menu-driven command-line application developed in **Go (Golang)** that demonstrates fundamental cybersecurity concepts in a practical manner. The tool integrates **classical cryptography** and **file integrity verification** into a single interactive application.

This project is intended primarily for **educational and academic purposes**, enabling students to understand how encryption and integrity mechanisms operate at an implementation level.

### Key Features

* Encrypt and decrypt messages using the **Rail Fence Cipher**
* Generate **SHA-256 hashes** for files
* Verify file integrity to detect modification, corruption, or tampering
* Simple, portable, and dependency-free design
=======
﻿#🛡️ Security Toolkit App (Go)
##✨ Short Description

The Security Toolkit App is a menu-driven command-line application built in Go (Golang) that demonstrates core cybersecurity concepts. It combines cryptography and file integrity verification into a single interactive tool.

The application allows users to:
>>>>>>> 2fb1eb4 (Final submission)

Encrypt and decrypt messages using the Rail Fence Cipher

<<<<<<< HEAD
## 📦 Dependencies and Libraries

This project relies **exclusively on the Go Standard Library**, ensuring ease of setup and cross-platform compatibility.

| Library         | Purpose                                       |
| --------------- | --------------------------------------------- |
| `fmt`           | Display menus, messages, and formatted output |
| `bufio`         | Read user input from the command line         |
| `os`            | File handling and system interaction          |
| `strings`       | String manipulation and input processing      |
| `strconv`       | Convert user input (e.g., cipher key)         |
| `crypto/sha256` | Generate secure SHA-256 hashes                |
| `encoding/hex`  | Convert hash bytes to hexadecimal format      |
| `io`            | Efficient file streaming for hashing          |
=======
Generate and verify SHA-256 hashes to detect file modification, corruption, or tampering

This project is designed for educational and academic purposes, helping students understand how encryption and integrity mechanisms work in practice.

##📦 Dependencies or Libraries Used
>>>>>>> 2fb1eb4 (Final submission)

This project uses only the Go Standard Library, ensuring simplicity, portability, and no external dependencies.

<<<<<<< HEAD
## 🛠️ Installation and Setup
=======
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
>>>>>>> 2fb1eb4 (Final submission)

Go (Golang)
Make sure Go is installed:

<<<<<<< HEAD
* **Go (Golang)** installed on your system

Verify installation:

```bash
go version
```

* Supported operating systems: **Windows, Linux, macOS**

### Setup Instructions

1. **Clone or download the repository**

   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the project directory**

   ```bash
   cd security-toolkit-app
   ```

3. **Run the application**

   ```bash
   go run *.go
   ```
=======
go version


Operating System
Works on Windows, Linux, and macOS.
>>>>>>> 2fb1eb4 (Final submission)

Setup

<<<<<<< HEAD
## 🚀 Usage

When the application starts, the following menu is displayed:

```
+----------------------------------+
|        Security Toolkit App       |
+----------------------------------+
| 1. Cipher Module                 |
| 2. File Integrity Checker        |
| 0. Exit                          |
+----------------------------------+
Choose an option:
```
=======
Clone or Download the Project

git clone <repository-url>


Navigate to the Project Directory

cd security-toolkit-app

>>>>>>> 2fb1eb4 (Final submission)

Run the Application

<<<<<<< HEAD
## 🔐 Cipher Module (Rail Fence Cipher)

### Encrypt a Message

**Steps:**

1. Select option **1** from the main menu
2. Choose **Encrypt Message**
3. Enter the plaintext message
4. Enter the number of rails (key)

**Example:**

```
Enter message: HELLO WORLD
Enter key (number of rails): 3
Encrypted Message: HOREL OLLWD
```

### Decrypt a Message

**Example:**

```
Enter message: HOREL OLLWD
Enter key (number of rails): 3
Decrypted Message: HELLO WORLD
```

> Note: The Rail Fence Cipher is a classical transposition cipher and is not secure for real-world applications.
=======
cd 

go run main.go encrypt.go decrypt.go verif.go hash.go utils.go 
>>>>>>> 2fb1eb4 (Final submission)

##🚀 Usage Examples

<<<<<<< HEAD
## 📁 File Integrity Checker

### Generate SHA-256 Hash

**Steps:**

1. Select option **2** from the main menu
2. Choose **Generate SHA-256 Hash**
3. Enter the file path

**Example:**

```
Enter file path: report.txt
SHA-256 Hash: a7c3f1b2e9...
```

### Verify File Integrity

**Example (File Intact):**

```
Enter file path: report.txt
Enter known hash: a7c3f1b2e9...
✅ File is intact
```

**Example (File Modified):**

```
❌ File has been modified
```
=======
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
>>>>>>> 2fb1eb4 (Final submission)

Steps:

<<<<<<< HEAD
## 🔄 Program Flow (High-Level)

```
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
```
=======
Choose option 1

Select Encrypt Message
>>>>>>> 2fb1eb4 (Final submission)

Enter the message

<<<<<<< HEAD
## 🎯 Learning Outcomes

By completing and using this project, users will:

* Understand **classical encryption techniques** (Rail Fence Cipher)
* Learn how **file integrity** is verified using cryptographic hash functions
* Practice building **menu-driven CLI applications** in Go
* Gain hands-on experience with **security-oriented programming concepts**

---

## 📌 Conclusion

The **Security Toolkit App** offers a concise and practical introduction to essential cybersecurity principles. By combining message encryption and file integrity verification into a single tool, the project effectively bridges theoretical knowledge and real-world implementation, making it well-suited for cybersecurity students and Go programming beginners.
=======
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
>>>>>>> 2fb1eb4 (Final submission)
