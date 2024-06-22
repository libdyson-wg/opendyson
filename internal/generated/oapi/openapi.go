// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
