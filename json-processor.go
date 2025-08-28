package main

import (
	"encoding/json"
	"os"
	"strconv"
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

func openFile(filepath string) []byte {
	var m LogMessage
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
