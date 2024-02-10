package builder

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/emp1re/goverter-test/xtype"
)

// ItemToListRule handles edge conditions if the target type is a pointer.
type Agg struct{}

// Matches returns true, if the builder can create handle the given types.
func (*Agg) Matches(_ *MethodContext, source, target *xtype.Type) bool {
	return !source.List && !source.Pointer && target.List
}

// Build creates conversion source code for the given source and target type.
func (*Agg) Build(gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type) ([]jen.Code, *xtype.JenID, *Error) {
	name := ctx.Name(target.ID())
	ctx.SetErrorTargetVar(jen.Nil())

	// Create a map of elements
	//mapElem := jen.Map(jen.Int()).Add(target.TypeAsJen()).Values()
	//
	//// Create a slice of elements
	//slice := jen.Index().Add(target.TypeAsJen()).Values()
	fmt.Println(name)
	stmt := append([]jen.Code{},
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Add(sourceID.Code),
			jen.If(jen.List(jen.Id("_"), jen.Id("ok")).Op(":=").Id("v.ID").Index(jen.Id("v.ID")), jen.Op("!ok"),
				jen.Id(name).Index(jen.Id("v.ID")).Op("=").Add(target.TypeAsJen()).Values(jen.Dict{
					jen.Id(name): jen.Id("v.ID"),
					jen.Id(name): jen.Id("v.Name"),
					jen.Id(name): jen.Index().String().Values(),
				}),
			),
			jen.Id("obj").Op(":=").Id(name).Index(jen.Id("v.ID")),
			jen.Id("obj.Addresses").Op("=").Append(jen.Id("obj.Addresses"), jen.Id("v.Address")),
			jen.Id(name).Index(jen.Id("v.ID")).Op("=").Id("obj"),
		),
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Add(sourceID.Code),
			jen.If(jen.List(jen.Id("found"), jen.Id("ok")).Op(":=").Id(name).Index(jen.Id("v.ID")), jen.Op("!ok"),
				jen.Continue(),
			),
			jen.Delete(jen.Id(name), jen.Id("v.ID")),
			jen.Id("result").Op("=").Append(jen.Id("result"), jen.Id("found")),
		),
	)

	newID := jen.Id(name)

	return stmt, xtype.OtherID(newID), nil
}
