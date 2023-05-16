package codegen

import (
	"errors"
	"github.com/Paulo-Lopes-Estevao/cli-sqlc-generate-sql-crud/internal/file"
)

type CodeGen struct {
	Content    []byte `default:""""`
	FileName   string
	PathTarget int
}

var (
	filePath   string
	existsBool bool
	files      = []string{"sqlc.yaml", "sqlc.yml", "sqlc.json"}
)

func ReadConfig() (string, error) {
	for _, f := range files {
		exists, err := file.CheckFileIfExists(f)
		if err != nil {
			return "", err
		}

		existsBool = exists

		if exists {
			filePath = f
			break
		}
	}

	if existsBool == false {
		return "", errors.New("error parsing configuration files (sqlc.yaml, sqlc.yml, sqlc.json) not found")
	}

	dir, err := file.GetPath(filePath)
	if err != nil {
		return "", err
	}

	filePath := file.JoinPath(dir, filePath)

	return filePath, nil
}

func GenerateCrudSql(data interface{}, pathTarget int, tag string) error {

	filePath, err := ReadConfig()
	if err != nil {
		return err
	}

	if filePath == "sqlc.yaml" || filePath == "sqlc.yml" {

		ymlConfig := NewYmlConfig()
		dataVersion, err := ymlConfig.ReadVerifyVersionYml(filePath)
		if err != nil {
			return err
		}

		err = dataVersion.GenerateCrudSqlYml(data, pathTarget, tag)
		if err != nil {
			return err
		}

	} else if filePath == "sqlc.json" {
		errors.New("developing...")
	}

	return nil
}
