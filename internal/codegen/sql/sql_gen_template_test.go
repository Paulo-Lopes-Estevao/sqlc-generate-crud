package sql

import (
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
