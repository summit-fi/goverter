// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import ignoremissing "github.com/summit-fi/goverter/example/ignore-missing"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source ignoremissing.Input) ignoremissing.Output {
	var exampleOutput ignoremissing.Output
	exampleOutput.Name = source.Name
	return exampleOutput
}
