package main

import (
	"encoding/json"
	"os"
)

// Harmless replacement for Aiven's discovery program. The hosted payload exits
// before the workflow's jq stage, so an empty successful discovery is sufficient.
func main() {
	if len(os.Args) >= 3 && os.Args[1] == "tests" && os.Args[2] == "matrix" {
		_ = json.NewEncoder(os.Stdout).Encode(map[string]any{
			"normal": []any{},
			"slow":   []any{},
		})
	}
}
