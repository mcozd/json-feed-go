package jsonfeed

import (
	"bytes"
	"encoding/json"
	"io"
)

type Parser struct{}

func FromString(jsonStr string) (feed *Feed, err error) {
	jsonBytes := bytes.NewBufferString(jsonStr).Bytes()
	return FromBytes(jsonBytes)
}

func FromBytes(jsonBytes []byte) (feed *Feed, err error) {
	err = json.Unmarshal(jsonBytes, &feed)
	return
}

func FromReader(reader io.Reader) (feed *Feed, err error) {
	jsonBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return FromBytes(jsonBytes)
}
