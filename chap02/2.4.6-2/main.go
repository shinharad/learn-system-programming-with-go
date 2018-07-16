package main

import (
	"encoding/json"
	"os"
)

// JSONを整形してio.Writerに書き出す
func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
