// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import ignore "github.com/emp1re/goverter-test/example/ignore"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source ignore.Input) ignore.Output {
	var exampleOutput ignore.Output
	exampleOutput.Name = source.Name
	return exampleOutput
}
