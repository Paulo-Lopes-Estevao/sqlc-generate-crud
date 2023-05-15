package sql

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateCrudSql(t *testing.T) {

	expectedContent := `
	-- name: GetAuthor :one
	SELECT * FROM author
	WHERE id = $1 LIMIT 1;

	-- name: ListAuthors :many
	SELECT * FROM author;

	-- name: CreateAuthor :one
	INSERT INTO author (
	  id,
	  name,
	  bio
	) VALUES (
	  $1,
	  $2,
	  $3
	)
	RETURNING *;

	-- name: UpdateAuthor :one
	UPDATE author
	SET 
	  id = $1,
	  name = $2,
	  bio = $3
	WHERE id = $4
	RETURNING *;

	-- name: DeleteAuthor :exec
	DELETE FROM author
	WHERE id = $1;
`

	structInfo := StructInfo{
		Name: "Author",
		Fields: []Field{
			{Name: "ID", Type: "int64"},
			{Name: "Name", Type: "string"},
			{Name: "Bio", Type: "string"},
		},
	}

	content, err := ContentTemplateCrudSql(structInfo)
	assert.NoError(t, err)

	assert.Equal(t, expectedContent, string(content))
}

func TestContentTemplateCrudSql(t *testing.T) {
	// Create a mock StructInfo
	structInfo := StructInfo{
		Name: "Author",
		Fields: []Field{
			{Name: "ID", Type: "int64"},
			{Name: "Name", Type: "string"},
			{Name: "Bio", Type: "string"},
		},
	}

	// Generate the SQL content using the ContentTemplateCrudSql method
	content, err := ContentTemplateCrudSql(structInfo)
	assert.NoError(t, err)

	tempFile, err := os.CreateTemp("", "author.sql")
	if err != nil {
		t.Fatal("Failed to create temporary file:", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatal("Failed to remove temporary file:", err)
		}
	}(tempFile.Name())

	// Write the generated content to the temporary file
	err = os.WriteFile(tempFile.Name(), content, 0644)
	assert.NoError(t, err)

	// Read the content of the "author.sql" file
	expectedContent, err := os.ReadFile(tempFile.Name())
	assert.NoError(t, err)

	// Compare the generated content with the content of the "author.sql" file
	assert.Equal(t, expectedContent, content)
}

func TestGetStructInfo(t *testing.T) {
	type Author struct {
		ID   int64  `bson:"id"`
		Name string `bson:"name"`
		Bio  string `bson:"bio"`
	}

	authorStructInfo := StructInfo{
		Name: "Author",
		Fields: []Field{
			{Name: "ID", Type: "int64", Format: "id"},
			{Name: "Name", Type: "string", Format: "name"},
			{Name: "Bio", Type: "string", Format: "bio"},
		},
	}

	author := Author{}
	structInfo, err := GetStructInfo(author, "bson")

	assert.NoError(t, err)
	assert.Equal(t, authorStructInfo, structInfo)
}
