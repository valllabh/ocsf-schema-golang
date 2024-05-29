package ocsfjson

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

type BaseEvent interface {
	GetActivityName() string
	GetCategoryName() string
	GetClassName() string
	GetCount() int32
	GetDuration() int32
	GetTime() int64
	GetMessage() string
	GetSeverity() string
	GetStatusCode() string
	GetTypeName() string

	proto.Message
}

func Marshal(m BaseEvent) ([]byte, error) {
	return json.Marshal(m)
}
