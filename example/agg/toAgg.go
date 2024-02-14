package agg

// goverter:converter
type Converter interface {
	Convert(Input1) Output1

	// goverter:agg ID Address
	ConvertAgg(source []Input1) []Output1
}
