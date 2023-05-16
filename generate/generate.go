package generate

import (
	"github.com/Paulo-Lopes-Estevao/cli-sqlc-generate-sql-crud/internal/codegen"
)

func Generate(data interface{}, pathTarget int, tag string) error {
	err := codegen.GenerateCrudSql(data, pathTarget, tag)
	if err != nil {
		return err
	}
	return nil
}
