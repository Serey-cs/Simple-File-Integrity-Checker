package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n+----------------------------------+")
		fmt.Println("|        Security Toolkit App       |")
		fmt.Println("+----------------------------------+")
		fmt.Println("| 1. Cipher Module                 |")
		fmt.Println("| 2. File Integrity Checker        |")
		fmt.Println("| 0. Exit                          |")
		fmt.Println("+----------------------------------+")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		// ================= Cipher Module =================
		case "1":
			fmt.Println("\n--- Cipher Module ---")
			fmt.Println("1. Encrypt Message")
			fmt.Println("2. Decrypt Message")
			fmt.Print("Choose an option: ")

			subChoice, _ := reader.ReadString('\n')
			subChoice = strings.TrimSpace(subChoice)

			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			fmt.Print("Enter key (number of rails): ")
			keyStr, _ := reader.ReadString('\n')
			keyStr = strings.TrimSpace(keyStr)
			key, _ := strconv.Atoi(keyStr)

			if subChoice == "1" {
				result := encryptRailFence(message, key)
				fmt.Println("Encrypted Message:", result)
			} else if subChoice == "2" {
				result := decryptRailFence(message, key)
				fmt.Println("Decrypted Message:", result)
			} else {
				fmt.Println("Invalid option.")
			}

		// ================= File Integrity Checker =================
		case "2":
			fmt.Println("\n--- File Integrity Checker ---")
			fmt.Println("1. Generate SHA-256 Hash")
			fmt.Println("2. Verify File Integrity")
			fmt.Print("Choose an option: ")

			subChoice, _ := reader.ReadString('\n')
			subChoice = strings.TrimSpace(subChoice)

			fmt.Print("Enter file path: ")
			filePath, _ := reader.ReadString('\n')
			filePath = strings.TrimSpace(filePath)

			if subChoice == "1" {
				hash, err := computeSHA256(filePath)
				if err != nil {
					fmt.Println("Error:", err)
					break
				}
				fmt.Println("SHA-256 Hash:")
				fmt.Println(hash)

			} else if subChoice == "2" {
				fmt.Print("Enter known hash: ")
				knownHash, _ := reader.ReadString('\n')
				knownHash = strings.TrimSpace(knownHash)

				match, err := verifyFile(filePath, knownHash)
				if err != nil {
					fmt.Println("Error:", err)
					break
				}

				if match {
					fmt.Println("✅ File is intact")
				} else {
					fmt.Println("❌ File has been modified")
				}
			} else {
				fmt.Println("Invalid option.")
			}

		case "0":
			fmt.Println("Exiting Security Toolkit. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
