package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		exitf("usage: %s <mode>\nallowed modes: false, store", filepath.Base(os.Args[0]))
	}

	mode := os.Args[1]
	extVal, err := modeToExtensionValue(mode)
	if err != nil {
		exitf("%v\nusage: %s <mode>\nallowed modes: false, store", err, filepath.Base(os.Args[0]))
	}

	const jsonPath = "custom/pocketstore.json"

	b, err := os.ReadFile(jsonPath)
	if err != nil {
		exitf("failed to read %s: %v", jsonPath, err)
	}

	// Use a generic map so we can preserve unknown fields.
	var doc map[string]any
	if err := json.Unmarshal(b, &doc); err != nil {
		exitf("failed to parse %s as JSON: %v", jsonPath, err)
	}

	doc["extension"] = extVal

	out, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		exitf("failed to serialize JSON: %v", err)
	}
	out = append(out, '\n')

	if err := os.WriteFile(jsonPath, out, 0644); err != nil {
		exitf("failed to write %s: %v", jsonPath, err)
	}

	fmt.Printf("updated %s: extension=%s\n", jsonPath, mode)
}

func modeToExtensionValue(mode string) (any, error) {
	switch mode {
	case "false":
		return false, nil
	case "store":
		return "store", nil
	default:
		return nil, errors.New("invalid mode: " + mode)
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}