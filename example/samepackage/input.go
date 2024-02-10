package samepackage

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package github.com/emp1re/goverter-test/example/samepackage
type Converter interface {
	Convert(source *Input) *Output
}

type Input struct{ Name string }
type Output struct{ Name string }
