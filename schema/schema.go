package schema

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GetSchema returns the GraphQL schema as a string
func GetSchema(schemaPath string) string {
	dat, err := ioutil.ReadFile(schemaPath)

	if err != nil {
		panic(err)
	}

	return string(dat[:])
}

// ConcatSchema concatenates all the schema files matching the glob
func ConcatSchema(schemaGlob string) string {
	m, err := filepath.Glob(schemaGlob)

	if err != nil {
		panic(err)
	}

	var sb strings.Builder
	for _, path := range m {
		sb.WriteString(GetSchema(path))
	}

	return sb.String()
}
