package validate

import (
	"embed"

	"github.com/valllabh/ocsf-schema-golang/ocsfvalidate/cache"
)

//go:embed jsonschema/*.json.zst
var schemaFS embed.FS

func Validate(className string, data []byte) error {
	return cache.SchemaValidator(
		schemaFS, className, "1.2", data,
	)
}
