// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package main

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
	"FDHJuvnu+5YmEkR2mdHivwEAAP//evDEExU9AAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
