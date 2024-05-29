package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"sync"

	"github.com/klauspost/compress/zstd"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Schema struct {
	className string
	version   string

	init    sync.Once
	fs      fs.FS
	js      *jsonschema.Schema
	loadErr error
}

var validatorCache sync.Map

func SchemaValidator(
	fs fs.FS,
	className string,
	version string,
	data []byte,
) error {
	ls, _ := validatorCache.LoadOrStore(
		className+":"+version,
		&Schema{className: className, version: version, fs: fs},
	)
	sch := ls.(*Schema)
	return sch.Validate(data)
}

func (s *Schema) load() (*jsonschema.Schema, error) {
	s.init.Do(func() {
		fn := fmt.Sprintf("jsonschema/classes_%s.json.zst", s.className)
		fh, err := s.fs.Open(fn)
		if err != nil {
			s.loadErr = err
			return
		}
		dec, err := zstd.NewReader(fh)
		if err != nil {
			s.loadErr = err
			return
		}
		buf := &bytes.Buffer{}
		_, err = io.Copy(buf, dec)
		if err != nil {
			s.loadErr = err
			return
		}

		sch, err := jsonschema.CompileString(s.className, buf.String())
		if err != nil {
			s.loadErr = err
			return
		}

		s.js = sch
	})
	return s.js, s.loadErr
}

func (s *Schema) Validate(data []byte) error {
	sch, err := s.load()
	if err != nil {
		return err
	}
	var v interface{}
	jd := json.NewDecoder(bytes.NewReader(data))
	jd.UseNumber()
	err = jd.Decode(&v)
	if err != nil {
		return err
	}

	return sch.Validate(v)
}
