package codegen

type CodeGen struct {
	Content    []byte `default:"nil"`
	FileName   string
	PathTarget int
}
