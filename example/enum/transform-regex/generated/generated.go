// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	"fmt"
	transformregex "github.com/jmattheis/goverter/example/enum/transform-regex"
)

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source transformregex.InputColor) transformregex.OutputColor {
	var exampleOutputColor transformregex.OutputColor
	switch source {
	case transformregex.ColBlue:
		exampleOutputColor = transformregex.BlueColor
	case transformregex.ColGreen:
		exampleOutputColor = transformregex.GreenColor
	case transformregex.ColRed:
		exampleOutputColor = transformregex.RedColor
	default:
		panic(fmt.Sprintf("unexpected enum element: %v", source))
	}
	return exampleOutputColor
}
func (c *ConverterImpl) Convert2(source transformregex.OutputColor) transformregex.InputColor {
	var exampleInputColor transformregex.InputColor
	switch source {
	case transformregex.BlueColor:
		exampleInputColor = transformregex.ColBlue
	case transformregex.GreenColor:
		exampleInputColor = transformregex.ColGreen
	case transformregex.RedColor:
		exampleInputColor = transformregex.ColRed
	default:
		panic(fmt.Sprintf("unexpected enum element: %v", source))
	}
	return exampleInputColor
}
