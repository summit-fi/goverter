// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import default1 "github.com/emp1re/goverter-test/example/default"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source *default1.Input) *default1.Output {
	pExampleOutput := default1.NewOutput()
	if source != nil {
		var exampleOutput default1.Output
		exampleOutput.Age = (*source).Age
		var pString *string
		if (*source).Name != nil {
			xstring := *(*source).Name
			pString = &xstring
		}
		exampleOutput.Name = pString
		pExampleOutput = &exampleOutput
	}
	return pExampleOutput
}
