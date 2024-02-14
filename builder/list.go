package builder

import (
	"github.com/dave/jennifer/jen"
	"github.com/summit-fi/goverter/xtype"
)

// List handles array / slice types.
type List struct{}

// Matches returns true, if the builder can create handle the given types.
func (*List) Matches(_ *MethodContext, source, target *xtype.Type) bool {
	return source.List && target.List && !target.ListFixed
}

// Build creates conversion source code for the given source and target type.
func (*List) Build(gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type) ([]jen.Code, *xtype.JenID, *Error) {

	agg, sFType, tFType := ctx.HasAggregation(ctx.Conf, source, target)
	if agg == true {

		// Set the error target variable to nil in the context
		ctx.SetErrorTargetVar(jen.Nil())
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
					if source.ListInner.StructType.Field(i).Name() == tFType.String {
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
			jen.Var().Id("result").Add(target.TypeAsJen()),
			// This line declares a map "m" with the key type of sFType (source field type) and value type of target.ListInner.String.
			jen.Id("m").Op(":=").Map(sFType.TypeAsJen()).Add(jen.Add(target.ListInner.TypeAsJen())).Values(),
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
	} else {
		ctx.SetErrorTargetVar(jen.Nil())
		targetSlice := ctx.Name(target.ID())
		index := ctx.Index()

		indexedSource := xtype.VariableID(sourceID.Code.Clone().Index(jen.Id(index)))

		errWrapper := Wrap("error setting index %d", jen.Id(index))
		forBlock, newID, err := gen.Build(ctx, indexedSource, source.ListInner, target.ListInner, errWrapper)
		if err != nil {
			return nil, nil, err.Lift(&Path{
				SourceID:   "[]",
				SourceType: source.ListInner.String,
				TargetID:   "[]",
				TargetType: target.ListInner.String,
			})
		}
		forBlock = append(forBlock, jen.Id(targetSlice).Index(jen.Id(index)).Op("=").Add(newID.Code))
		forStmt := jen.For(jen.Id(index).Op(":=").Lit(0), jen.Id(index).Op("<").Len(sourceID.Code.Clone()), jen.Id(index).Op("++")).
			Block(forBlock...)

		stmt := []jen.Code{}
		if source.ListFixed {
			stmt = []jen.Code{
				jen.Id(targetSlice).Op(":=").Make(target.TypeAsJen(), jen.Len(sourceID.Code.Clone())),
				forStmt,
			}
		} else {
			stmt = []jen.Code{
				jen.Var().Add(jen.Id(targetSlice), target.TypeAsJen()),
				jen.If(sourceID.Code.Clone().Op("!=").Nil()).Block(
					jen.Id(targetSlice).Op("=").Make(target.TypeAsJen(), jen.Len(sourceID.Code.Clone())),
					forStmt,
				),
			}
		}

		return stmt, xtype.VariableID(jen.Id(targetSlice)), nil
	}

}
