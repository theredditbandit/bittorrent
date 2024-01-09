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
	} else if command == "info" {
		fileName := os.Args[2]
		fileReader, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		content, err := bencode.Decode(fileReader)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		metaInfo, _ := content.(map[string]interface{})
        fmt.Printf("Tracker URL: %v\n",metaInfo["announce"])
        info , _ := metaInfo["info"].(map[string]interface{})
        fmt.Printf("Length: %v\n",info["length"])

	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
