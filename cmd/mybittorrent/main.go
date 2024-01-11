package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jackpal/bencode-go"
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
		fmt.Printf("Info Hash: %x\n", torrentFile.getInfoHash())
        fmt.Printf("Piece Length: %v\n", torrentFile.Info.PieceLength)
        torrentFile.printPieceHashes()
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
