# 🛡️ Security Toolkit App (Go)

## ✨ Overview

The **Security Toolkit App** is a menu-driven command-line application developed in **Go (Golang)** that demonstrates fundamental cybersecurity concepts in a practical manner. The tool integrates **classical cryptography** and **file integrity verification** into a single interactive application.

This project is intended primarily for **educational and academic purposes**, enabling students to understand how encryption and integrity mechanisms operate at an implementation level.

### Key Features

* Encrypt and decrypt messages using the **Rail Fence Cipher**
* Generate **SHA-256 hashes** for files
* Verify file integrity to detect modification, corruption, or tampering
* Simple, portable, and dependency-free design

---

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

---

## 🛠️ Installation and Setup

### Prerequisites

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

---

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

---

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

---

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

---

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

---

## 🎯 Learning Outcomes

By completing and using this project, users will:

* Understand **classical encryption techniques** (Rail Fence Cipher)
* Learn how **file integrity** is verified using cryptographic hash functions
* Practice building **menu-driven CLI applications** in Go
* Gain hands-on experience with **security-oriented programming concepts**

---

## 📌 Conclusion

The **Security Toolkit App** offers a concise and practical introduction to essential cybersecurity principles. By combining message encryption and file integrity verification into a single tool, the project effectively bridges theoretical knowledge and real-world implementation, making it well-suited for cybersecurity students and Go programming beginners.
