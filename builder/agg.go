package builder

import (
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/emp1re/goverter-test/xtype"
)

// Agg handles edge conditions if the target type is a pointer.
type Agg struct{}

// Matches returns true, if the builder can create handle the given types.
func (*Agg) Matches(_ *MethodContext, source, target *xtype.Type) bool {
	return source.List && !source.Pointer && target.List
}

// Build creates conversion source code for the given source and target type.
func (*Agg) Build(gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type) ([]jen.Code, *xtype.JenID, *Error) {

	// Set the error target variable to nil in the context
	ctx.SetErrorTargetVar(jen.Nil())

	// Split the raw field settings from the context by space
	// This is done to extract individual settings from a single string.
	split := strings.Split(ctx.Conf.RawFieldSettings[0], " ")

	// Check if the split result has less than 3 elements.
	// If it does, return nil for all return values.
	// This is because we expect at least 3 elements from the split operation.
	// The first element is ignored, and the second and third elements are used.
	if len(split) < 3 {
		return nil, nil, &Error{
			Cause: "Invalid settings for goverter:agg",
		}
	}

	// Extract the second element from the split result and assign it to the variable "mark".
	// This represents the first relevant setting from the raw field settings.
	mark := split[1]

	// Extract the third element from the split result and assign it to the variable "mark2".
	// This represents the second relevant setting from the raw field settings.
	mark2 := split[2]

	// Initialize an xtype.Type for the source
	var sFType xtype.Type
	// Iterate over the fields of the source struct type
	for i := 0; i < source.ListInner.StructType.NumFields(); i++ {
		// If the field name matches "mark", set the xtype.Type to the field information
		if mark == source.ListInner.StructType.Field(i).Name() {
			sFType = xtype.Type{
				String: source.ListInner.StructType.Field(i).Name(),
				T:      source.ListInner.StructType.Field(i).Type(),
			}
		}
	}

	// Initialize an xtype.Type for the target
	var tFType xtype.Type

	// Iterate over the fields of the target struct type
	for i := 0; i < target.ListInner.StructType.NumFields(); i++ {
		// If the field name matches "mark2", set the xtype.Type to the field information
		if mark2 == target.ListInner.StructType.Field(i).Name() {
			tFType = xtype.Type{
				String: target.ListInner.StructType.Field(i).Name(),
				T:      target.ListInner.StructType.Field(i).Type(),
			}

		}
	}
	// Declare a variable "empty" of type string. This variable will be used to hold
	// the "empty" value for the target field type.
	var empty string
	switch tFType.T.String() {
	case "[]string":
		empty = "\"\""
	case "[]int":
		empty = "0"
	case "[]int8":
		empty = "0"
	case "[]int16":
		empty = "0"
	case "[]int32":
		empty = "0"
	case "[]int64":
		empty = "0"
	case "[]uint":
		empty = "0"
	case "[]uint8":
		empty = "0"
	case "[]uint16":
		empty = "0"
	case "[]uint32":
		empty = "0"
	case "[]uint64":
		empty = "0"
	case "[]float":
		empty = "0.0"
	case "[]float32":
		empty = "0.0"
	case "[]float64":
		empty = "0.0"
	case "[]complex64":
		empty = "0.0"
	case "[]complex128":
		empty = "0.0"
	case "[]bool":
		empty = "false"
	case "[]rune":
		empty = "0"
	default:
		return nil, nil, &Error{
			Cause: "Invalid settings for goverter:agg -> expected different type for target field",
		}

	}

	// blockIter is a slice of jen.Code. It represents a block of code that is generated dynamically.
	// This block of code is used to create a new instance of the target struct type.
	// The fields of the new instance are populated based on the fields of the source struct type.
	blockIter := []jen.Code{
		// jen.Id(target.ListInner.String).ValuesFunc creates a new instance of the target struct type.
		// The function passed to ValuesFunc is used to populate the fields of the new instance.
		jen.Id(target.ListInner.String).BlockFunc(func(g *jen.Group) {
			// This loop iterates over the fields of the source struct type.
			for i := 0; i < source.ListInner.StructType.NumFields(); i++ {
				// If the name of the current field matches mark2, a new slice of strings is created and assigned to the field.
				if source.ListInner.StructType.Field(i).Name() == mark2 {
					g.Id(source.ListInner.StructType.Field(i).Name()).Op(":").Id(tFType.T.String()).Values().Op(",")
					continue
				}
				// If the name of the current field does not match mark2, the value of the field in the source struct is assigned to the corresponding field in the target struct.
				g.Id(target.ListInner.StructType.Field(i).Name()).Op(":").Id("v").Dot(source.ListInner.StructType.Field(i).Name()).Op(",")
			}
		}),
	}
	// resultBlock is a slice of jen.Code that represents a block of code for handling the result of the conversion.
	resultBlock := []jen.Code{
		// This line tries to get the value from the map "m" with the key "v.sFType".
		// "found" is the value found in the map, and "ok" is a boolean indicating whether the key was found in the map.
		jen.List(jen.Id("found"), jen.Id("ok")).Op(":=").Id("m").Index(jen.Id("v").Dot(sFType.String)),
		// This conditional statement checks if the key was not found in the map. If the key was not found, it continues to the next iteration of the loop.
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Continue(),
		),
		// This line deletes the key "v.sFType" from the map "m".
		jen.Delete(jen.Id("m"), jen.Id("v").Dot(sFType.String)),
		// This line appends the value found in the map to the "result" variable.
		jen.Id("result").Op("=").Append(jen.Id("result"), jen.Id("found")),
	}
	// Define the block of code
	block := []jen.Code{
		// This line attempts to retrieve a value from the map "m" using the key "v.sFType".
		// The underscore "_" is a blank identifier, used when syntax requires a variable name but program logic does not.
		// "ok" is a boolean that is true if the key was found in the map, and false otherwise.
		jen.If(
			jen.List(jen.Id("_"), jen.Id("ok")).Op(":=").Id("m").Index(jen.Id("v").Dot(sFType.String)),
			jen.Id("!ok"),
		).Block(
			// If the key was not found in the map ("ok" is false), this line sets the value in the map with the key "v.sFType" to the target type.
			// The target type is represented by the blockIter slice of jen.Code.
			jen.Id("m").Index(jen.Id("v").Dot(sFType.String)).Op("=").Add(blockIter...),
		),
		jen.If(jen.Id("v").Dot(tFType.String)).Op("==").Id(empty).Block(
			jen.Continue()),
		// This line retrieves the object from the map "m" using the key "v.sFType" and assigns it to the variable "obj".
		jen.Id("obj").Op(":=").Id("m").Index(jen.Id("v").Dot(sFType.String)),
		// This line appends the string "v.tFType.String" to the slice "obj.tFType.String".
		jen.Id("obj").Dot(tFType.String).Op("=").Append(jen.Id("obj").Dot(tFType.String), jen.Id("v").Dot(tFType.String)),
		// This line sets the value in the map with the key "v.sFType" to the object "obj".
		jen.Id("m").Index(jen.Id("v").Dot(sFType.String)).Op("=").Id("obj"),
	}

	// Define the SMTP block of code
	smtp := []jen.Code{
		// This line declares a variable "result" with the type of the target.
		jen.Var().Id("result").Id(target.String),
		// This line declares a map "m" with the key type of sFType (source field type) and value type of target.ListInner.String.
		jen.Id("m").Op(":=").Map(sFType.TypeAsJen()).Add(jen.Id(target.ListInner.String)).Values(),
		// This loop iterates over the source. For each iteration, it executes the block of code defined earlier.
		// The underscore "_" is a blank identifier, used when syntax requires a variable name but program logic does not.
		// "v" is the value of the current iteration.
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("source")).Block(block...),
		// This loop iterates over the source again. For each iteration, it executes the resultBlock of code defined earlier.
		// The underscore "_" is a blank identifier, used when syntax requires a variable name but program logic does not.
		// "v" is the value of the current iteration.
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("source")).Block(resultBlock...),
	}

	// Return the SMTP block of code, the result variable identifier, and nil for error
	return smtp, xtype.VariableID(jen.Id("result")), nil
}
