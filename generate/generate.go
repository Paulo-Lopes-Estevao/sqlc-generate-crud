// Package generate provides functionality for generating SQL CRUD operations.
package generate

import (
	"github.com/Paulo-Lopes-Estevao/sqlc-generate-crud/internal/codegen"
)

// Generate generates SQL CRUD operations based on the provided data and configuration.
// The generated code will be saved to the specified target path.
// The tag parameter allows specifying a struct tag for the generated code.
// If an error occurs during code generation, it will be returned.
func Generate(data interface{}, pathTarget int, tag string) error {
	err := codegen.GenerateCrudSql(data, pathTarget, tag)
	if err != nil {
		return err
	}
	return nil
}
