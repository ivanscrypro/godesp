package main

import (
	"encoding/json"
	"errors"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Payloads struct {
	Mode []struct {
		Day   []Methods `json:"day"`
		Night []Methods `json:"night"`
	} `json:"mode"`
}

type Methods struct {
	Get []struct {
		Routes     []string `json:"routes"`
		Parameters []string `json:"parameters"`
		Payloads   []string `json:"payloads"`
	} `json:"GET"`
	Post []struct {
		Routes     []string `json:"routes"`
		Parameters []string `json:"parameters"`
		Payloads   []string `json:"payloads"`
	} `json:"POST"`
	Options []struct {
		Routes     []string `json:"routes"`
		Parameters []string `json:"parameters"`
		Payloads   []string `json:"payloads"`
	} `json:"OPTIONS"`
	Patch []struct {
		Routes     []string `json:"routes"`
		Parameters []string `json:"parameters"`
		Payloads   []string `json:"payloads"`
	} `json:"PATCH"`
}

// Validate a *single* file name (no separators) you might store/display.
// Use in addition to SafeJoin if you accept bare names from users.
func ValidateFileName(name string) error {
	var (
		ErrBadName = errors.New("invalid file name")
		ErrNotUTF8 = errors.New("name not valid UTF-8")
	)
	if name == "" {
		return ErrBadName
	}
	if !utf8.ValidString(name) {
		return ErrNotUTF8
	}
	// Forbid path separators and control chars.
	if strings.ContainsRune(name, '/') || strings.ContainsRune(name, '\\') {
		return ErrBadName
	}
	for _, r := range name {
		if r < 0x20 { // control chars
			return ErrBadName
		}
	}
	// Forbid device names on Windows (NUL, CON, PRN, AUX, COM1, LPT1, etc.)
	if runtime.GOOS == "windows" {
		l := strings.ToLower(name)
		bad := []string{"con", "prn", "aux", "nul",
			"com1", "com2", "com3", "com4", "com5", "com6", "com7", "com8", "com9",
			"lpt1", "lpt2", "lpt3", "lpt4", "lpt5", "lpt6", "lpt7", "lpt8", "lpt9"}
		for _, b := range bad {
			if l == b || strings.HasPrefix(l, b+".") {
				return ErrBadName
			}
		}
	}
	return nil
}

func openFile(filepath string) []byte {
	var m LogMessage
	err := ValidateFileName(filepath)
	if err != nil {
		m.MessageType = "fatal"
		m.Message = "The filename " + filepath + " is invalid or contains invalid characters"
		m.getLogger()
		return nil
	}
	// reading the device-keywords.json file
	content, err := os.ReadFile(filepath)
	if err != nil {
		m.MessageType = "fatal"
		m.Message = "There is no file " + filepath
		m.getLogger()
		return nil
	}
	return content
}

func (Payloads *Payloads) readJSON(assets string) {
	var m LogMessage
	content := openFile(assets)
	err := json.Unmarshal(content, &Payloads)
	if err != nil || Payloads.Mode == nil {
		m.MessageType = "regular"
		m.Message = "Check if the file " + assets + " is in json format"
		m.getLogger()
	}
}

func (Payloads *Payloads) writeJSON(assets string) error {
	var m LogMessage
	output, err := json.MarshalIndent(Payloads, "", "  ")
	if err != nil {
		m.MessageType = "error"
		m.Message = "Can not Marshal the Payloads to JSON format"
		m.getLogger()
		return err
	}
	err = os.WriteFile(assets, output, 0666)
	if err != nil {
		m.MessageType = "error"
		m.Message = "Can not write to the " + assets + " file"
		m.getLogger()
		return err
	}
	return nil
}

func dividePayload(assets string) {
	var (
		m LogMessage
		p Payloads
	)
	// Defining numbers to test the 'count' vs them
	methods := 4
	for i := 0; i < methods; i++ {
		p.readJSON(assets)
		if p.Mode == nil {
			m.MessageType = "regular"
			m.Message = "JSON to divide is empty"
			m.getLogger()
			os.Exit(1)
		}
		switch i {
		case 0:
			savePayloads(i, "get", p)
		case 1:
			savePayloads(i, "post", p)
		case 2:
			savePayloads(i, "options", p)
		case 3:
			savePayloads(i, "patch", p)
		}

	}
}

func savePayloads(i int, method string, payload Payloads) {
	switch method {
	case "get":
		payload.Mode[0].Day[0].Post = nil
		payload.Mode[0].Day[0].Options = nil
		payload.Mode[0].Day[0].Patch = nil
	case "post":
		payload.Mode[0].Day[0].Get = nil
		payload.Mode[0].Day[0].Options = nil
		payload.Mode[0].Day[0].Patch = nil
	case "options":
		payload.Mode[0].Day[0].Get = nil
		payload.Mode[0].Day[0].Post = nil
		payload.Mode[0].Day[0].Patch = nil
	case "patch":
		payload.Mode[0].Day[0].Get = nil
		payload.Mode[0].Day[0].Post = nil
		payload.Mode[0].Day[0].Options = nil
	}
	payload.writeJSON("./assets/payloads" + strconv.Itoa(i) + ".json")
}
