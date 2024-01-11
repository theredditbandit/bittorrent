package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/jackpal/bencode-go"
)

type metaInfo struct {
	Announce string
	Info     struct {
		Length      int `bencode:"length"`
		Name        string `bencode:"name"`
		PieceLength int `bencode:"piece length"`
		Pieces      string `bencode:"pieces"`
	}
}

func (m metaInfo) printPieceHashes() {
    pieces := []byte(m.Info.Pieces)
    fmt.Println("Piece Hashes:")
    for i := 0; i< len(pieces); i += 20 {
        hash := pieces[i:i+20]
        fmt.Printf("%x\n",hash)
    }
}

// hashes the Info struct
func (m metaInfo) getInfoHash() string {
	bencodedString := m.encode()
	h := sha1.New()
	h.Write([]byte(bencodedString))
	return string(h.Sum(nil))
}

// encodes the Info struct as bencode
func (m metaInfo) encode() string {
	s := bytes.NewBufferString("")
	err := bencode.Marshal(s, m.Info)
	if err != nil {
		fmt.Printf("err encode(): %v\n", err)
	}
	return s.String()
}

func (m metaInfo) decode(s string) interface{} {
	buf := bytes.NewBufferString(s)
	decoded, _ := bencode.Decode(buf)
	return decoded
}
