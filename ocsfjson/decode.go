package ocsfjson

import "google.golang.org/protobuf/encoding/protojson"

func Unmarshal(b []byte, m BaseEvent) error {
	return protojson.Unmarshal(b, m)
}
