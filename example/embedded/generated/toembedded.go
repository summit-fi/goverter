// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import embedded "github.com/emp1re/goverter-test/example/embedded"

type ToConverterImpl struct{}

func (c *ToConverterImpl) ToEmbedded(source embedded.FlatPerson) embedded.Person {
	var examplePerson embedded.Person
	examplePerson.Address = c.ToEmbeddedAddress(source)
	examplePerson.Name = source.Name
	return examplePerson
}
func (c *ToConverterImpl) ToEmbeddedAddress(source embedded.FlatPerson) embedded.Address {
	var exampleAddress embedded.Address
	exampleAddress.Street = source.StreetName
	exampleAddress.ZipCode = source.ZipCode
	return exampleAddress
}
