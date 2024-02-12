// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import extendlocalwithconverter "github.com/emp1re/goverter-test/example/extend-local-with-converter"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source extendlocalwithconverter.Input) extendlocalwithconverter.Output {
	var exampleOutput extendlocalwithconverter.Output
	exampleOutput.Animals = extendlocalwithconverter.ConvertAnimals(c, source.Animals)
	return exampleOutput
}
func (c *ConverterImpl) ConvertCats(source []extendlocalwithconverter.Cat) []extendlocalwithconverter.Animal {
	var exampleAnimalList []extendlocalwithconverter.Animal
	if source != nil {
		exampleAnimalList = make([]extendlocalwithconverter.Animal, len(source))
		for i := 0; i < len(source); i++ {
			exampleAnimalList[i] = c.exampleCatToExampleAnimal(source[i])
		}
	}
	return exampleAnimalList
}
func (c *ConverterImpl) ConvertDogs(source []extendlocalwithconverter.Dog) []extendlocalwithconverter.Animal {
	var exampleAnimalList []extendlocalwithconverter.Animal
	if source != nil {
		exampleAnimalList = make([]extendlocalwithconverter.Animal, len(source))
		for i := 0; i < len(source); i++ {
			exampleAnimalList[i] = c.exampleDogToExampleAnimal(source[i])
		}
	}
	return exampleAnimalList
}
func (c *ConverterImpl) exampleCatToExampleAnimal(source extendlocalwithconverter.Cat) extendlocalwithconverter.Animal {
	var exampleAnimal extendlocalwithconverter.Animal
	exampleAnimal.Name = source.Name
	return exampleAnimal
}
func (c *ConverterImpl) exampleDogToExampleAnimal(source extendlocalwithconverter.Dog) extendlocalwithconverter.Animal {
	var exampleAnimal extendlocalwithconverter.Animal
	exampleAnimal.Name = source.Name
	return exampleAnimal
}
