package main

import (
	"fmt"
	"github.com/jackpal/bencode-go"
	"strings"
)

func decodeBencode(bencodedString string) (interface{}, error) {
	myreader := strings.NewReader(bencodedString)
	data, err := bencode.Decode(myreader)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return data, nil
}
