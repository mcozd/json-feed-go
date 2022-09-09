package jsonfeed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Parser struct{}

func FromString(jsonStr string) (feed *Feed, err error) {
	jsonBytes := bytes.NewBufferString(jsonStr).Bytes()
	return FromBytes(jsonBytes)
}

func FromBytes(jsonBytes []byte) (feed *Feed, err error) {
	err = json.Unmarshal(jsonBytes, &feed)
	if err == nil && (feed == nil || feed.IsEmpty()) {
		err = fmt.Errorf("unmarshalling ok, but Feed object is nil or empty, check the json for valid jsonfeed data")
	}
	return
}

func FromReader(reader io.Reader) (feed *Feed, err error) {
	jsonBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return FromBytes(jsonBytes)
}
