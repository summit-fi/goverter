// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import skipcopysametype "github.com/summit-fi/goverter/example/skip-copy-same-type"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source skipcopysametype.Input) skipcopysametype.Output {
	var exampleOutput skipcopysametype.Output
	exampleOutput.Name = source.Name
	exampleOutput.ItemCounts = source.ItemCounts
	return exampleOutput
}
