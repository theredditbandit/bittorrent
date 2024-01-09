package main

import (
	"encoding/json"
	"fmt"
	"github.com/jackpal/bencode-go"
	"os"
	"strings"
)

func main() {
	command := os.Args[1]
	if command == "decode" {
		bencodedValue := os.Args[2]
		stringReader := strings.NewReader(bencodedValue)
		decoded, err := bencode.Decode(stringReader)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
