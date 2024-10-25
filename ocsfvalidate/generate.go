//go:build gen_json_schema

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/klauspost/compress/zstd"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func main() {
	versions := []string{"1.0", "1.1", "1.2"}

	for _, version := range versions {
		classlist := fmt.Sprintf(
			"../schema/v%s-schema-class-list.json",
			version,
		)

		d, err := os.ReadFile(classlist)
		if err != nil {
			log.Fatalf("reading %s: %w", classlist, err)
			return
		}

		cl := map[string]interface{}{}
		err = json.Unmarshal(d, &cl)
		if err != nil {
			log.Fatalf("parsing schema-class-list.json: %w", err)
			return
		}
		for _, v := range cl {
			vv := v.([]interface{})
			for _, vvx := range vv {
				className := vvx.(string)
				log.Printf("fetching %s/%s", className, version)
				err = fetchSchema(className, version)
				if err != nil {
					log.Fatalf("failed fetching %s: %w", className, err)
					return
				}
			}
		}
	}
}

func fetchSchema(className string, version string) error {
	fullVersion := fmt.Sprintf("%s.0", version)

	ux := url.URL{
		Scheme: "https",
		Host:   "schema.ocsf.io",
		Path:   path.Join("schema", fullVersion, "classes", className),
		RawQuery: url.Values{
			"profiles": []string{""},
		}.Encode(),
	}

	// 	https://schema.ocsf.io/schema/1.2.0/classes/web_resources_activity?profiles=
	diskPath := filepath.Join(version, "jsonschema", "classes_"+className+".json.zst")
	dpAbs, err := filepath.Abs(diskPath)
	if err != nil {
		return fmt.Errorf("fetchSchema: filepath.Abs failed: %w", err)
	}
	diskURL := url.URL{
		Scheme: "file",
		Path:   filepath.ToSlash(dpAbs),
	}

	req, err := http.NewRequest("GET", ux.String(), nil)
	if err != nil {
		return fmt.Errorf("fetchSchema: http.NewRequest failed: %w", err)
	}
	req.Header.Set("User-Agent", "ocsf-schema-golang/fetch-schema")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("fetchSchema: http fetch failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(
		http.MaxBytesReader(nil, resp.Body, 1024*1024*10),
	)
	if err != nil {
		return fmt.Errorf("fetchSchema: http read all failed: %w", err)
	}
	buf := &bytes.Buffer{}
	err = json.Indent(buf, data, "", "  ")
	if err != nil {
		return fmt.Errorf("fetchSchema: json Ident failed: %w", err)
	}

	_, err = jsonschema.CompileString(diskURL.String(), buf.String())
	if err != nil {
		return fmt.Errorf("fetchSchema: jsonschema.CompileString failed: %w", err)
	}

	out := &bytes.Buffer{}
	enc, err := zstd.NewWriter(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(enc, buf)
	if err != nil {
		enc.Close()
		return err
	}
	err = enc.Close()
	if err != nil {
		return err
	}

	err = os.WriteFile(diskPath, out.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("fetchSchema: write file failed: %w", err)
	}

	return nil
}
