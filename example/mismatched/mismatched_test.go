package mismatched_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/summit-fi/goverter/example/mismatched"
	"github.com/summit-fi/goverter/example/mismatched/generated"
)

func TestConverter(t *testing.T) {
	var c mismatched.Converter = &generated.ConverterImpl{}

	input := mismatched.DBCustomers{
		{
			DBPerson: mismatched.DBPerson{
				First: "mary",
				Last:  "brown",
			},
		},
		{
			DBPerson: mismatched.DBPerson{
				First: "john",
				Last:  "smith",
			},
			DBAddress: &mismatched.DBAddress{
				Suburb:   "Surry Hills",
				Postcode: "2010",
			},
		},
	}

	actual := c.Convert(input)

	expected := mismatched.APICustomers{
		{
			APIPerson: &mismatched.APIPerson{
				First: "mary",
				Last:  "brown",
			},
		},
		{
			APIPerson: &mismatched.APIPerson{
				First: "john",
				Last:  "smith",
			},
			APIAddress: mismatched.APIAddress{
				Suburb:   "Surry Hills",
				Postcode: "2010",
			},
		},
	}

	bytes, err := json.MarshalIndent(expected, "", " ")
	if err != nil {
		panic(err)
	}
	t.Log(string(bytes))
	require.Equal(t, expected, actual)
}
