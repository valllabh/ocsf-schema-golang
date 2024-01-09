package enumValueMap

import (
	"embed"
	"encoding/json"
)

//go:embed enum-value-map.json
var embedFs embed.FS

// struct to hold the enum value
type EnumValue struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type EnumValueMap = map[string]EnumValue

var enumValueMap EnumValueMap

func init() {
	data, _ := embedFs.ReadFile("enum-value-map.json")
	json.Unmarshal(data, &enumValueMap)
}

func GetEnumValue(name string) (EnumValue, bool) {
	enumValue, enumValueExits := enumValueMap[name]
	return enumValue, enumValueExits
}