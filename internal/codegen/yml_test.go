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

func TestReadYmlVersion1(t *testing.T) {
	content := []byte(`version: "1"
packages:
  - name: "db"
    path: "internal/db"
    queries: "./sql/query/"
    schema: "./sql/schema/"
    engine: "postgresql"
`)

	// Call the method
	ymlVersion1, err := ReadYmlVersion1(content)

	// Check for errors
	assert.NoError(t, err, "ReadYmlVersion1 should not return an error")

	// Validate the result
	expectedPackage := YmlVersion1{
		Package: []struct {
			Name    string `yaml:"name"`
			Engine  string `yaml:"engine"`
			Queries string `yaml:"queries"`
			Schema  string `yaml:"schema"`
		}{
			{
				Name:    "db",
				Engine:  "postgresql",
				Queries: "./sql/query/",
				Schema:  "./sql/schema/",
			},
		},
	}

	assert.Equal(t, expectedPackage, *ymlVersion1, "Incorrect package")

	// Alternatively, you can validate individual fields separately
	assert.Len(t, ymlVersion1.Package, 1, "Incorrect number of packages")
	assert.Equal(t, "db", ymlVersion1.Package[0].Name, "Incorrect package name")
	assert.Equal(t, "postgresql", ymlVersion1.Package[0].Engine, "Incorrect engine")
	assert.Equal(t, "./sql/query/", ymlVersion1.Package[0].Queries, "Incorrect queries path")
	assert.Equal(t, "./sql/schema/", ymlVersion1.Package[0].Schema, "Incorrect schema path")
}

func TestReadYmlVersion2(t *testing.T) {

	yamlContent := []byte(`version: "2"
sql:
  - schema: "postgresql/schema.sql"
    queries: "postgresql/query.sql"
    engine: "postgresql"`)

	ymlVersion2, err := ReadYmlVersion2(yamlContent)

	assert.NoError(t, err, "ReadYmlVersion2 should not return an error")

	expectedPackage := YmlVersion2{
		SQL: []struct {
			Engine  string `yaml:"engine"`
			Queries string `yaml:"queries"`
			Schema  string `yaml:"schema"`
		}{
			{
				Engine:  "postgresql",
				Queries: "postgresql/query.sql",
				Schema:  "postgresql/schema.sql",
			},
		},
	}

	assert.Equal(t, expectedPackage, *ymlVersion2, "Incorrect package")

	// Alternatively, you can validate individual fields separately
	assert.Len(t, ymlVersion2.SQL, 1, "Incorrect number of packages")
	assert.Equal(t, "postgresql", ymlVersion2.SQL[0].Engine, "Incorrect engine")
	assert.Equal(t, "postgresql/query.sql", ymlVersion2.SQL[0].Queries, "Incorrect queries path")
	assert.Equal(t, "postgresql/schema.sql", ymlVersion2.SQL[0].Schema, "Incorrect schema path")

}
