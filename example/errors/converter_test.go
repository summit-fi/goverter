package errors_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/summit-fi/goverter/example/errors"
	"github.com/summit-fi/goverter/example/errors/generated"
)

func TestConverterSuccess(t *testing.T) {
	var c errors.Converter = &generated.ConverterImpl{}

	input := errors.APIApartment{
		Position: 5,
		Owner: errors.APIPerson{
			ID:       0,
			FullName: "j mattheis",
		},
	}

	temp, err := c.ToDBApartment(input)
	require.NoError(t, err)
	actual := c.ToAPIApartment(temp)

	require.Equal(t, input, actual)
}

func TestConverterError(t *testing.T) {
	var c errors.Converter = &generated.ConverterImpl{}

	input := errors.APIApartment{
		Position: 5,
		Owner: errors.APIPerson{
			ID:       0,
			FullName: "oops",
		},
	}

	_, err := c.ToDBApartment(input)
	require.Error(t, err)
}
