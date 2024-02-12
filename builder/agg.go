package builder

import (
	"regexp"
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

	// Get the current index from the context
	index := ctx.Index()
	// Create a new variable identifier with the name "v"
	indexedSource := xtype.VariableID(jen.Id("v"))

	// Wrap the error message with the current index
	errWrapper := Wrap("error setting index %d", jen.Id(index))
	// Build the conversion source code for the given source and target type
	_, id, err := gen.Build(ctx, indexedSource, source.ListInner, target.ListInner, errWrapper)
	// If there is an error, lift it with the path information and return
	if err != nil {
		return nil, nil, err.Lift(&Path{
			SourceID:   "[]",
			SourceType: source.ListInner.String,
			TargetID:   "[]",
			TargetType: target.ListInner.String,
		})
	}

	// Set the error target variable to nil in the context
	ctx.SetErrorTargetVar(jen.Nil())
	// Compile a regular expression to match "[]"
	re := regexp.MustCompile(`\[\]`)
	// Replace all occurrences of "[]" in the target string with ""
	str := re.ReplaceAllString(target.String, "")

	// Split the raw field settings from the context by space
	split := strings.Split(ctx.Conf.RawFieldSettings[0], " ")
	// Get the second and third elements from the split result
	mark := split[1]
	mark2 := split[2]
	// Define the result block of code
	resulBlock := []jen.Code{
		// Try to get the value from the map with the key "v.mark"
		jen.List(jen.Id("found"), jen.Id("ok")).Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
		// If the key is not found, continue to the next iteration
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Continue(),
		),
		// Delete the key "v.mark" from the map
		jen.Delete(jen.Id("i"), jen.Id("v").Dot(mark)),
		// Append the found value to the result
		jen.Id("result").Op("=").Append(jen.Id("result"), jen.Id("found")),
	}
	// Define the block of code
	block := []jen.Code{
		// Try to get the value from the map with the key "v.mark"
		jen.If(
			jen.List(jen.Id("_"), jen.Id("ok")).Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
			jen.Id("!ok"),
		).Block(
			// If the key is not found, set the value to the generated code
			jen.Id("i").Index(jen.Id("v").Dot(mark)).Op("=").Add(id.Code),
			// Get the object from the map with the key "v.mark"
			jen.Id("obj").Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
			// Set the "obj.mark2" to an empty string slice
			jen.Id("obj").Dot(mark2).Op("=").Index().String().Values(),
			// Set the value in the map with the key "v.mark" to the object
			jen.Id("i").Index(jen.Id("v").Dot(mark)).Op("=").Id("obj"),
		),
		// Get the object from the map with the key "v.mark"
		jen.Id("obj").Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
		// Append "v.mark2" to "obj.mark2"
		jen.Id("obj").Dot(mark2).Op("=").Append(jen.Id("obj").Dot(mark2), jen.Id("v").Dot(mark2)),
		// Set the value in the map with the key "v.mark" to the object
		jen.Id("i").Index(jen.Id("v").Dot(mark)).Op("=").Id("obj"),
	}

	// Initialize an xtype.Type for the source
	var xType xtype.Type
	// Iterate over the fields of the source struct type
	for i := 0; i < source.ListInner.StructType.NumFields(); i++ {
		// If the field name matches "mark", set the xtype.Type to the field information
		if mark == source.ListInner.StructType.Field(i).Name() {
			xType = xtype.Type{
				String: source.ListInner.StructType.Field(i).Name(),
				T:      source.ListInner.StructType.Field(i).Type(),
			}
		}
	}

	//// Initialize an xtype.Type for the target
	//var targeType xtype.Type
	//
	//// Iterate over the fields of the target struct type
	//for i := 0; i < target.ListInner.StructType.NumFields(); i++ {
	//	// If the field name matches "mark2", set the xtype.Type to the field information
	//	if mark2 == target.ListInner.StructType.Field(i).Name() {
	//		targeType = xtype.Type{
	//			String: target.ListInner.StructType.Field(i).Name(),
	//			T:      target.ListInner.StructType.Field(i).Type(),
	//		}
	//	}
	//}

	// Define the SMTP block of code
	smtp := []jen.Code{
		// Declare a variable "result" with the type of the target
		jen.Var().Id("result").Id(target.String),
		// Declare a map "i" with the key type of xType and value type of str
		jen.Id("i").Op(":=").Map(xType.TypeAsJen()).Add(jen.Id(str)).Values(),
		// Iterate over the source
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("source")).Block(block...),
		// Iterate over the source again
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("source")).Block(resulBlock...),
	}

	// Return the SMTP block of code, the result variable identifier, and nil for error
	return smtp, xtype.VariableID(jen.Id("result")), nil
}
