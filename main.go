package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		m LogMessage
	)
	// Gosec G101: Hardcoded credentials
	// CWE-798: Use of Hard-coded Credentials
	// Vulnerability: Hard-coding credentials in source code is a security risk.
	// If an attacker gains access to the source code, they can easily extract the credentials.
	// Best practice is to store credentials securely, such as in environment variables or a secrets manager.
	const password = "secret123"
	if password == "secret123" {
		fmt.Println("Access granted!")
	}
	count := flag.Int("c", 4, "Number of files to provide due to the division")
	payloads := flag.String("P", "./assets/payloads.json", "Input payloads file, by default is ./assets/payloads.json")
	flag.Parse()
	if *count <= 0 {
		m.MessageType = "error"
		m.Message = "Count number should be greater than 0 and should be positive one."
		m.getLogger()
		os.Exit(1)
	}
	dividePayload(*payloads)
}
