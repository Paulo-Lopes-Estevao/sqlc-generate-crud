package codegen

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadVerifyVersionYml(t *testing.T) {
	tempFile, err := os.CreateTemp("", "sqlc.yaml")
	if err != nil {
		t.Fatal("Failed to create temporary file:", err)
	}
	defer os.Remove(tempFile.Name())

	yamlContent := []byte(`version: "2"
sql:
  - schema: "postgresql/schema.sql"
    queries: "postgresql/query.sql"
    engine: "postgresql"`)
	err = os.WriteFile(tempFile.Name(), yamlContent, 0644)
	if err != nil {
		t.Fatal("Failed to write YAML content to the temporary file:", err)
	}

	ymlConfig := NewYmlConfig()

	result, err := ymlConfig.ReadVerifyVersionYml(tempFile.Name())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "2", result.Version)
	assert.NotNil(t, result.YmlVersion2)
	assert.Nil(t, result.YmlVersion1)
	assert.Equal(t, 1, len(result.YmlVersion2.SQL))
	assert.Equal(t, "postgresql", result.YmlVersion2.SQL[0].Engine)
	assert.Equal(t, "postgresql/query.sql", result.YmlVersion2.SQL[0].Queries)
	assert.Equal(t, "postgresql/schema.sql", result.YmlVersion2.SQL[0].Schema)
}
