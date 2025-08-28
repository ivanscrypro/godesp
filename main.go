package main

import (
	"flag"
	"os"
)

func main() {
	var (
		m LogMessage
	)
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
