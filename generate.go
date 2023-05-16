// Package sqlcgeneratecrud provides a set of utilities for generating CRUD SQL statements such as INSERT, SELECT, UPDATE and DELETE.
//
// Source code and other details for the project are available at GitHub:
//
// https://github.com/Paulo-Lopes-Estevao/sqlc-generate-crud
package sqlcgeneratecrud

import "github.com/Paulo-Lopes-Estevao/sqlc-generate-crud/internal/codegen"

// GenerateConfig represents the configuration options for generating CRUD SQL statements.
type GenerateConfig struct {
	TagFormat  string // The tag format to use for struct field tags (default: "json").
	PathTarget int    // The target path for the generated SQL files (default: 0).
}

// Generate generates CRUD SQL statements for the given data using the specified options.
// It returns an error if any error occurs during the generation process.
func Generate(data interface{}, option *GenerateConfig) error {
	err := codegen.GenerateCrudSql(data, option.pathTarget(), option.tag())
	if err != nil {
		return err
	}
	return nil
}

// tag returns the tag format to use. If not specified, it defaults to "json".
func (g *GenerateConfig) tag() string {
	if g.TagFormat == "" {
		g.TagFormat = "json"
	}
	return g.TagFormat
}

// pathTarget returns the target path for the generated SQL files.
func (g *GenerateConfig) pathTarget() int {
	return g.PathTarget
}
