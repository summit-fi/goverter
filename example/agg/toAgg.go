package agg

// goverter:converter
type Converter interface {
	Convert(Input1) Output1

	// goverter:agg Address Address
	ConvertAgg(source []Input1) []Output1
}
