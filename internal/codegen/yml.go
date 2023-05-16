package codegen

import (
	"errors"
	"fmt"
	"github.com/Paulo-Lopes-Estevao/cli-sqlc-generate-sql-crud/internal/codegen/sql"
	"github.com/Paulo-Lopes-Estevao/cli-sqlc-generate-sql-crud/internal/file"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"strings"
)

type IYmlConfig interface {
	ReadVerifyVersionYml(filePath string) (*YmlConfig, error)
	GenerateCrudSqlYml(data interface{}, pathTarget int, tag string) error
	WriteCrudSql() error
	CheckDirectoryQueriesValuePathYmlV1() error
	CreateDirectoryQueriesSQLV1() error
	CheckDirectoryQueriesValuePathYmlV2() error
	CreateDirectoryQueriesSQLV2() error
}

type YmlConfig struct {
	Version     string `yaml:"version"`
	CodeGen     CodeGen
	YmlVersion1 *YmlVersion1
	YmlVersion2 *YmlVersion2
}

type YmlVersion1 struct {
	Package []struct {
		Name    string `yaml:"name"`
		Engine  string `yaml:"engine"`
		Queries string `yaml:"queries"`
		Schema  string `yaml:"schema"`
	} `yaml:"packages"`
}

type YmlVersion2 struct {
	SQL []struct {
		Engine  string `yaml:"engine"`
		Queries string `yaml:"queries"`
		Schema  string `yaml:"schema"`
	} `yaml:"sql"`
}

func NewYmlConfig() *YmlConfig {
	return &YmlConfig{}
}

func (ymlConfig *YmlConfig) ReadVerifyVersionYml(filePath string) (*YmlConfig, error) {

	content, err := file.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &ymlConfig)
	if err != nil {
		return nil, err
	}

	if ymlConfig.Version == "1" {

		ymlConfig.YmlVersion1, err = ReadYmlVersion1(content)
		if err != nil {
			return nil, err
		}
		return ymlConfig, nil

	} else if ymlConfig.Version == "2" {

		ymlConfig.YmlVersion2, err = ReadYmlVersion2(content)
		if err != nil {
			return nil, err
		}
		return ymlConfig, nil

	} else {
		return nil, errors.New("version not found")
	}

}

func (ymlConfig *YmlConfig) GenerateCrudSqlYml(data interface{}, pathTarget int, tag string) error {
	structInfo, err := sql.GetStructInfo(data, tag)
	if err != nil {
		return err
	}

	content, err := sql.ContentTemplateCrudSql(structInfo)
	if err != nil {
		return err
	}

	ymlConfig.CodeGen.Content = content
	ymlConfig.CodeGen.FileName = strings.ToLower(structInfo.Name)
	ymlConfig.CodeGen.PathTarget = pathTarget

	if err := ymlConfig.WriteCrudSql(); err != nil {
		return err
	}

	return nil
}

func ReadYmlVersion1(content []byte) (*YmlVersion1, error) {
	var ymlVersion1 YmlVersion1

	err := yaml.Unmarshal(content, &ymlVersion1)
	if err != nil {
		return nil, err
	}

	return &ymlVersion1, nil
}

func ReadYmlVersion2(content []byte) (*YmlVersion2, error) {
	var ymlVersion2 YmlVersion2

	err := yaml.Unmarshal(content, &ymlVersion2)
	if err != nil {
		return nil, err
	}

	return &ymlVersion2, nil
}

func (ymlConfig *YmlConfig) WriteCrudSql() error {
	if ymlConfig.Version == "1" {
		return ymlConfig.CheckDirectoryQueriesValuePathYmlV1()
	} else if ymlConfig.Version == "2" {
		return ymlConfig.CheckDirectoryQueriesValuePathYmlV2()
	} else {
		return errors.New("version not found")
	}
}

func (ymlConfig *YmlConfig) CheckDirectoryQueriesValuePathYmlV2() error {
	if len(ymlConfig.YmlVersion2.SQL) >= 0 {
		if err := ymlConfig.CreateDirectoryQueriesSQLV2(); err != nil {
			return err
		}
	} else {
		return errors.New("not found path queries")
	}

	return nil
}

func (ymlConfig *YmlConfig) CreateDirectoryQueriesSQLV2() error {

	dir := filepath.Dir(ymlConfig.YmlVersion2.SQL[ymlConfig.CodeGen.PathTarget].Queries)

	filePath := file.JoinPath(dir, fmt.Sprintf("%s.sql", ymlConfig.CodeGen.FileName))

	val, err := file.CheckFileIfExists(filePath)
	if err != nil {
		return err
	}
	if !val {
		if err := file.CreateDirAndFile(filePath, ymlConfig.CodeGen.Content); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("file %s already exists", filePath)
	}

	return nil
}

func (ymlConfig *YmlConfig) CheckDirectoryQueriesValuePathYmlV1() error {
	if len(ymlConfig.YmlVersion1.Package) >= 0 {
		if err := ymlConfig.CreateDirectoryQueriesSQLV1(); err != nil {
			return err
		}
	} else {
		return errors.New("not found path queries")
	}

	return nil
}

func (ymlConfig *YmlConfig) CreateDirectoryQueriesSQLV1() error {
	dir := filepath.Dir(ymlConfig.YmlVersion1.Package[ymlConfig.CodeGen.PathTarget].Queries)

	filePath := file.JoinPath(dir, fmt.Sprintf("%s.sql", ymlConfig.CodeGen.FileName))

	val, err := file.CheckFileIfExists(filePath)
	if err != nil {
		return err
	}
	if !val {
		if err := file.CreateDirAndFile(filePath, ymlConfig.CodeGen.Content); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("file %s already exists", filePath)
	}

	return nil
}
