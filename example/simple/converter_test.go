package simple_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/summit-fi/goverter/example/simple"
	"github.com/summit-fi/goverter/example/simple/generated"
)

func TestConverter(t *testing.T) {
	var c simple.Converter = &generated.ConverterImpl{}
	inputs := []simple.Input{
		{Age: 5, Name: "jmattheis"},
		{Age: 75, Name: "other"},
	}

	actual := c.Convert(inputs)

	expected := []simple.Output{
		{Age: 5, Name: "jmattheis"},
		{Age: 75, Name: "other"},
	}

	require.Equal(t, expected, actual)
}
