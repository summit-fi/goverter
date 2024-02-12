package builder

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/emp1re/goverter-test/xtype"
)

// ItemToListRule handles edge conditions if the target type is a pointer.
type Agg struct{}

// Matches returns true, if the builder can create handle the given types.
func (*Agg) Matches(_ *MethodContext, source, target *xtype.Type) bool {
	return source.List && !source.Pointer && target.List
}

// Build creates conversion source code for the given source and target type.
func (*Agg) Build(gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type) ([]jen.Code, *xtype.JenID, *Error) {
	//name := ctx.Name(target.ID())
	//ctx.SetErrorTargetVar(jen.Nil())

	// Create a map of elements
	//mapElem := jen.Map(jen.Int()).Add(target.TypeAsJen()).Values()
	//
	//// Create a slice of elements
	//slice := jen.Index().Add(target.TypeAsJen()).Values()

	//stmt := append([]jen.Code{},
	//	jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Add(sourceID.Code),
	//		jen.If(jen.List(jen.Id("_"), jen.Id("ok")).Op(":=").Id("v.ID").Index(jen.Id("v.ID")), jen.Op("!ok"),
	//			jen.Id(name).Index(jen.Id("v.ID")).Op("=").Add(target.TypeAsJen()).Values(jen.Dict{
	//				jen.Id(name): jen.Id("v.ID"),
	//				jen.Id(name): jen.Id("v.Name"),
	//				jen.Id(name): jen.Index().String().Values(),
	//			}),
	//		),
	//		jen.Id("obj").Op(":=").Id(name).Index(jen.Id("v.ID")),
	//		jen.Id("obj.Addresses").Op("=").Append(jen.Id("obj.Addresses"), jen.Id("v.Address")),
	//		jen.Id(name).Index(jen.Id("v.ID")).Op("=").Id("obj"),
	//	),
	//	jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Add(sourceID.Code),
	//		jen.If(jen.List(jen.Id("found"), jen.Id("ok")).Op(":=").Id(name).Index(jen.Id("v.ID")), jen.Op("!ok"),
	//			jen.Continue(),
	//		),
	//		jen.Delete(jen.Id(name), jen.Id("v.ID")),
	//		jen.Id("result").Op("=").Append(jen.Id("result"), jen.Id("found")),
	//	),
	//)
	//
	//newID := jen.Id(name)
	//ctx.SetErrorTargetVar(jen.Nil())
	//targetSlice := ctx.Name(target.ID())

	//
	//errWrapper := Wrap("error setting index %d", jen.Id(index))
	//forBlock, newID, err := gen.Build(ctx, indexedSource, source.ListInner, target.ListInner, errWrapper)
	//if err != nil {
	//	return nil, nil, err.Lift(&Path{
	//		SourceID:   "[]",
	//		SourceType: source.ListInner.String,
	//		TargetID:   "[]",
	//		TargetType: target.ListInner.String,
	//	})
	//}
	//forBlock = append(forBlock, jen.Id(targetSlice).Index(jen.Id(index)).Op("=").Add(newID.Code))
	//forStmt := jen.For(jen.Id(index).Op(":=").Lit(0), jen.Id(index).Op("<").Len(sourceID.Code.Clone()), jen.Id(index).Op("++")).
	//	Block(forBlock...)
	//
	//stmt := []jen.Code{}
	//if source.ListFixed {
	//	stmt = []jen.Code{
	//		jen.Id(targetSlice).Op(":=").Make(target.TypeAsJen(), jen.Len(sourceID.Code.Clone())),
	//		forStmt,
	//	}
	//} else {
	//	stmt = []jen.Code{
	//		jen.Var().Add(jen.Id(targetSlice), target.TypeAsJen()),
	//		jen.If(sourceID.Code.Clone().Op("!=").Nil()).Block(
	//			jen.Id(targetSlice).Op("=").Make(target.TypeAsJen(), jen.Len(sourceID.Code.Clone())),
	//			forStmt,
	//		),
	//	}
	//}
	//
	index := ctx.Index()

	//indexedSource := xtype.VariableID(sourceID.Code.Clone().Index(jen.Id(index)))
	indexedSource := xtype.VariableID(jen.Id("v"))

	errWrapper := Wrap("error setting index %d", jen.Id(index))
	_, id, err := gen.Build(ctx, indexedSource, source.ListInner, target.ListInner, errWrapper)
	if err != nil {
		return nil, nil, err.Lift(&Path{
			SourceID:   "",
			SourceType: source.ListInner.String,
			TargetID:   "",
			TargetType: target.ListInner.String,
		})
	}

	//return stmt, xtype.VariableID(jen.Id(targetSlice)), nil
	name := ctx.Name(target.ID())
	sN := ctx.Name(source.ID())
	fmt.Println("#########################")

	fmt.Println("listInner", source.ListInner.StructType.String())

	fmt.Println("ctx ----->", ctx.Conf.Definition)
	fmt.Println("gen", gen)
	ctx.SetErrorTargetVar(jen.Nil())
	re := regexp.MustCompile(`\[\]`)
	str := re.ReplaceAllString(target.String, "")
	fmt.Println("source", str)
	fmt.Println("sourseID")
	fmt.Println("source", source)
	fmt.Println("target", target)
	fmt.Println("source.ID", source.ID())
	fmt.Println("target.ID", target.ID())
	fmt.Println("sourceID", sourceID)
	fmt.Println("target.ID", name)
	fmt.Println("source.ID", sN)
	fmt.Println("ctx--->", ctx.Conf)

	key, value := ctx.Map()
	fmt.Println("ctx.Map", key, value)

	targetSlice := ctx.Name(target.ID())

	fmt.Println("targetSlice", targetSlice)

	split := strings.Split(ctx.Conf.RawFieldSettings[0], " ")
	mark := split[1]
	mark2 := split[2]
	resulBlock := []jen.Code{
		jen.List(jen.Id("found"), jen.Id("ok")).Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
		jen.If(jen.Op("!").Id("ok")).Block(
			jen.Continue(),
		),
		jen.Delete(jen.Id("i"), jen.Id("v").Dot(mark)),
		jen.Id("result").Op("=").Append(jen.Id("result"), jen.Id("found")),
	}
	block := []jen.Code{
		jen.If(
			jen.List(jen.Id("_"), jen.Id("ok")).Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
			jen.Id("!ok"),
		).Block(
			jen.Id("i").Index(jen.Id("v").Dot(mark)).Op("=").Add(id.Code),
		),
		jen.Id("obj").Op(":=").Id("i").Index(jen.Id("v").Dot(mark)),
		jen.Id("obj").Dot(mark2).Op("=").Append(jen.Id("obj").Dot(mark2), jen.Id("v").Dot(mark2)),
		jen.Id("i").Index(jen.Id("v").Dot(mark)).Op("=").Id("obj"),
	}
	var xType xtype.Type
	//sourseStruct := split[2]
	for i := 0; i < source.ListInner.StructType.NumFields(); i++ {
		if mark == source.ListInner.StructType.Field(i).Name() {
			xType = xtype.Type{
				String: source.ListInner.StructType.Field(i).Name(),
				T:      source.ListInner.StructType.Field(i).Type(),
			}
		}
	}
	targetStruct := target.ListInner.StructType
	fmt.Println("targetStruct", targetStruct)
	fmt.Println(targetStruct.String())
	var targeType xtype.Type
	for i := 0; i < target.ListInner.StructType.NumFields(); i++ {
		if mark2 == target.ListInner.StructType.Field(i).Name() {
			targeType = xtype.Type{
				String: target.ListInner.StructType.Field(i).Name(),
				T:      target.ListInner.StructType.Field(i).Type(),
			}
		}
	}
	fmt.Println("targeType", targeType)
	fmt.Println(source.ListInner)
	list := source.ListInner.StructType.Field(0)
	fmt.Println("list", list.Type())

	//newVar jen.Code{
	//	jen.Var().Id(str).Add(target.TypeAsJen()),
	//}
	smtp := []jen.Code{
		jen.Var().Id("result").Id(target.String),
		//jen.Var().Id("obj").Id(target.ListInner.String),
		jen.Id("i").Op(":=").Map(xType.TypeAsJen()).Add(jen.Id(str)).Values(),
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("source")).Block(block...),
		jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id("source")).Block(resulBlock...),
	}
	//smtp = append(smtp, forBlock...)

	return smtp, xtype.VariableID(jen.Id("result")), nil

	//fmt.Println("target.ID", target.ID())
	//fmt.Println("target.ID().Code", target.ID().Code)
	//fmt.Println("target.ID().Code.Clone()", )

	//smtp := []jen.Code{
	//	jen.Id("m").Op(":=").Map(jen.Int()).Id(source.String).Values(),
	//
	//
	//	jen.For(jen.List(jen.Id("_"), jen.Id("v")).Op(":=").Range().Id(source.ID(),), jen.Block(
	//}
	//
	//mapID := jen.Id("m")
	//vID := jen.Id("v")
	//okID := jen.Id("ok")
	//objID := jen.Id("obj")
	//foundID := jen.Id("found")
	//resultID := jen.Id("result")
	//
	//stmt := []jen.Code{
	//	mapID.Op(":=").Map(jen.Int()).Id(source.String).Values(),
	//	jen.For(vID.Op(":=").Range().Id(source.String), jen.Block(
	//		jen.List(jen.Id("_"), okID).Op(":=").Id(source.String).Index(vID.Id("ID")),
	//		jen.If(okID).Block(
	//			jen.Id(source.String).Index(vID.Id("ID")).Op("=").Id("Output1").Values(jen.Dict{
	//				jen.Id("ID"):        vID.Id("ID"),
	//				jen.Id("Name"):      vID.Id("Name"),
	//				jen.Id("Addresses"): jen.Index().String().Values(),
	//			}),
	//		),
	//		objID.Op(":=").Id(source.String).Index(vID.Id("ID")),
	//		objID.Id("Addresses").Op("=").Append(objID.Id("Addresses"), vID.Id("Address")),
	//		jen.Id(source.String).Index(vID.Id("ID")).Op("=").Id(source.String),
	//	)),
	//	resultID.Op(":=").Index().Id("Output1").Values(),
	//	jen.For(vID.Op(":=").Range().Id(name), jen.Block(
	//		jen.List(foundID, okID).Op(":=").Id(source.String).Index(vID.Id("ID")),
	//		jen.If(okID).Block(
	//			jen.Delete(mapID, vID.Id("ID")),
	//			resultID.Op("=").Append(resultID, foundID),
	//		),
	//	)),
	//	jen.Return(resultID),
	//}
	//
	//return nil, nil, nil
	//return stmt, xtype.OtherID(newID), nil
}
