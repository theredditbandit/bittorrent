package main

import (
	"encoding/json"
	"fmt"
	"github.com/jackpal/bencode-go"
	"os"
	"strings"
)

type metaInfo struct {
	Announce string
	Info     struct {
		Length      int
		Name        string
		PieceLength int `bencode:"piece length"`
		Pieces      string
	}
}

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
		r, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("err file opening: %v\n", err)
		}
		var torrentFile metaInfo
		err = bencode.Unmarshal(r, &torrentFile)
		if err != nil {
			fmt.Printf("err unmarshalling: %v\n", err)
		}
        fmt.Printf("Tracker URL: %s\n", torrentFile.Announce)
        fmt.Printf("Length: %d\n", torrentFile.Info.Length)

	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
