package sql

import (
	"errors"
	"reflect"
	"strings"
	"text/template"
)

type Field struct {
	Name   string
	Type   string
	Format string
}

type StructInfo struct {
	Name   string
	Fields []Field
}

func TemplateCrudSql(structInfo StructInfo) (*template.Template, error) {
	const crudTemplate = `
	-- name: Get{{ .Name }} :one
	SELECT * FROM {{ .Name | ToSnakeCase }}
	WHERE id = $1 LIMIT 1;

	-- name: List{{ Pluralize .Name }} :many
	SELECT * FROM {{ .Name | ToSnakeCase }};

	-- name: Create{{ .Name }} :one
	INSERT INTO {{ .Name | ToSnakeCase }} (
	{{- range $index, $field := .Fields }}
	  {{ $field.Name | ToSnakeCase }}{{ if not (last $index (len $.Fields)) }},{{ end }}
	{{- end }}
	) VALUES (
	{{- range $index, $field := .Fields }}
	  ${{ add $index 1 }}{{ if not (last $index (len $.Fields)) }},{{ end }}
	{{- end }}
	)
	RETURNING *;

	-- name: Update{{ .Name }} :one
	UPDATE {{ .Name | ToSnakeCase }}
	SET {{ range $index, $field := .Fields }}
	  {{ $field.Name | ToSnakeCase }} = ${{ add $index 1 }}{{ if not (last $index (len $.Fields)) }},{{ end }}
	{{- end }}
	WHERE id = ${{ add (len .Fields) 1 }}
	RETURNING *;

	-- name: Delete{{ .Name }} :exec
	DELETE FROM {{ .Name | ToSnakeCase }}
	WHERE id = $1;
`
	tmpl := template.Must(template.New("crudTemplate").Funcs(template.FuncMap{
		"ToSnakeCase": ToSnakeCase,
		"Pluralize":   Pluralize,
		"add":         func(a, b int) int { return a + b },
		"last":        LastFunc,
	}).Parse(crudTemplate))

	return tmpl, nil
}

func GetStructInfo(data interface{}, tag string) (StructInfo, error) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Struct {
		return StructInfo{}, errors.New("input is not a struct")
	}

	structInfo := StructInfo{
		Name:   value.Type().Name(),
		Fields: make([]Field, 0, value.NumField()),
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		structInfo.Fields = append(structInfo.Fields, Field{
			Name:   field.Name,
			Type:   field.Type.Name(),
			Format: getFieldTagValue(field, tag),
		})
	}

	return structInfo, nil
}

func getFieldTagValue(field reflect.StructField, tag string) string {
	return strings.Split(field.Tag.Get(tag), ",")[0]
}

func ToSnakeCase(str string) string {
	var result strings.Builder

	for i, char := range str {
		if i > 0 && (char >= 'A' && char <= 'Z') && str[i-1] != 'I' && str[i-1] != 'D' {
			result.WriteByte('_')
		}
		result.WriteRune(char)
	}

	return strings.ToLower(result.String())
}

func Pluralize(str string) string {
	return str + "s"
}

func ContentTemplateCrudSql(structInfo StructInfo) ([]byte, error) {
	tmpl, err := TemplateCrudSql(structInfo)
	if err != nil {
		return nil, err
	}

	var content strings.Builder
	err = tmpl.Execute(&content, structInfo)
	if err != nil {
		return nil, err
	}

	return []byte(content.String()), nil
}

func LastFunc(index, length int) bool {
	return index == length-1
}
