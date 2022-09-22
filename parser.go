package jsonfeed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type FeedConstraint interface {
	IsEmpty() bool
}

func FromBytes[F FeedConstraint](jsonBytes []byte) (feed *F, err error) {

	err = json.Unmarshal(jsonBytes, &feed)
	if err == nil && (feed == nil || (*feed).IsEmpty()) {
		err = fmt.Errorf("unmarshalling ok, but Feed object is nil or empty, check the json for valid jsonfeed data")
	}
	return
}

func FromString[F FeedConstraint](jsonStr string) (feed *F, err error) {
	jsonBytes := bytes.NewBufferString(jsonStr).Bytes()
	return FromBytes[F](jsonBytes)
}

func FromReader[F FeedConstraint](reader io.Reader) (feed *F, err error) {
	if reader == nil {
		return nil, fmt.Errorf("reader was nil")
	}
	jsonBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return FromBytes[F](jsonBytes)
}
