package main

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// PSEUDOCODE — WON'T COMPILE

// 2) Command Injection (os/exec with untrusted arg)
func runTool(arg string) {
	// ❌ unvalidated user input into system command
	_ = exec.Command("sh", "-c", "convert "+arg+" /tmp/out.png") // placeholder
}

// 3) Path Traversal (user-controlled file path)
func readFile(p string) {
	// ❌ no normalization / allowlisting; e.g. "../../etc/passwd"
	_, _ = os.ReadFile(p)
}

// 4) Weak crypto / insecure config
func insecureTLS(c *http.Client) {
	c.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ❌
	}
}

func badHash(data []byte) []byte {
	// ❌ MD5 for security purpose
	h := md5.Sum(data)
	return h[:]
}

// 6) Untrusted deserialization / JSON into interface{} without checks
func parse(data []byte) {
	var v interface{}
	_ = json.Unmarshal(data, &v) // ❌ unchecked dynamic types used later
}

func hardcode() {
	// Gosec G101: Hardcoded credentials
	// CWE-798: Use of Hard-coded Credentials
	// Vulnerability: Hard-coding credentials in source code is a security risk.
	// If an attacker gains access to the source code, they can easily extract the credentials.
	// Best practice is to store credentials securely, such as in environment variables or a secrets manager.
	const password = "secret123"
	if password == "secret123" {
		fmt.Println("Access granted!")
	}
}
