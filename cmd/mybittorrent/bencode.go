package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	// "github.com/jackpal/bencode-go"
)

func decodeBencode(bencodedString string) (interface{}, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {
		return decodeStr(bencodedString)
	} else if strings.HasPrefix(bencodedString, "i") && strings.HasSuffix(bencodedString, "e") {
		return decodeInt(bencodedString)
	} else {
		return "", fmt.Errorf("only strings are supported at the moment")
	}
}

func decodeInt(bencodedString string) (interface{}, error) {
	bencodeint := bencodedString[1 : len(bencodedString)-1]
	decodedint, err := strconv.Atoi(bencodeint)
	if err != nil {
		return "", err
	}
	return decodedint, nil
}

func decodeStr(bencodedString string) (interface{}, error) {
	var firstColonIndex int
	for i := 0; i < len(bencodedString); i++ {
		if bencodedString[i] == ':' {
			firstColonIndex = i
			break
		}
	}
	lengthStr := bencodedString[:firstColonIndex]
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", err
	}
	return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], nil

}
