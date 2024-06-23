// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	common "github.com/jmattheis/goverter/example/format/common"
	interfacetostruct "github.com/jmattheis/goverter/example/format/interfacetostruct"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertApartment(source common.DBApartment) common.APIApartment {
	var commonAPIApartment common.APIApartment
	commonAPIApartment.Position = source.Position
	commonAPIApartment.Owner = c.ConvertPerson(source.Owner)
	commonAPIApartment.OwnerName = source.Owner.Name
	if source.CoResident != nil {
		commonAPIApartment.CoResident = make([]common.APIPerson, len(source.CoResident))
		for i := 0; i < len(source.CoResident); i++ {
			commonAPIApartment.CoResident[i] = c.ConvertPerson(source.CoResident[i])
		}
	}
	return commonAPIApartment
}
func (c *ConverterImpl) ConvertHouse(source common.DBHouse) common.APIHouse {
	var commonAPIHouse common.APIHouse
	commonAPIHouse.Address = source.Address
	commonAPIHouse.Apartments = interfacetostruct.ConvertToApartmentMap(c, source.Apartments)
	return commonAPIHouse
}
func (c *ConverterImpl) ConvertPerson(source common.DBPerson) common.APIPerson {
	var commonAPIPerson common.APIPerson
	commonAPIPerson.ID = source.ID
	commonAPIPerson.MiddleName = interfacetostruct.SQLStringToPString(source.MiddleName)
	pString := source.Name
	commonAPIPerson.FirstName = &pString
	if source.Friends != nil {
		commonAPIPerson.Friends = make([]common.APIPerson, len(source.Friends))
		for i := 0; i < len(source.Friends); i++ {
			commonAPIPerson.Friends[i] = c.ConvertPerson(source.Friends[i])
		}
	}
	commonAPIPerson.Info = c.commonInfoToCommonInfo(source.Info)
	return commonAPIPerson
}
func (c *ConverterImpl) commonInfoToCommonInfo(source common.Info) common.Info {
	var commonInfo common.Info
	commonInfo.Birthplace = source.Birthplace
	return commonInfo
}
