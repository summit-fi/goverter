package builder

import (
	"github.com/dave/jennifer/jen"

	"github.com/summit-fi/goverter/xtype"
)

// ItemToListRule handles edge conditions if the target type is a pointer.
type ItemToListRule struct{}

// Matches returns true, if the builder can create handle the given types.
func (*ItemToListRule) Matches(_ *MethodContext, source, target *xtype.Type) bool {
	return !source.List && !source.Pointer && target.List
}

// Build creates conversion source code for the given source and target type.
func (*ItemToListRule) Build(gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type) ([]jen.Code, *xtype.JenID, *Error) {
	name := ctx.Name(target.ID())
	ctx.SetErrorTargetVar(jen.Nil())

	// Create a slice of elements
	slice := jen.Index().Add(source.TypeAsJen()).ValuesFunc(func(g *jen.Group) {
		g.Add(sourceID.Code)
	})

	stmt := append([]jen.Code{}, jen.Id(name).Op(":=").Add(slice))
	newID := jen.Id(name)

	return stmt, xtype.OtherID(newID), nil
}
