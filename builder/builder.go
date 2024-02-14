package builder

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/summit-fi/goverter/config"
	"github.com/summit-fi/goverter/method"
	"github.com/summit-fi/goverter/namer"
	"github.com/summit-fi/goverter/xtype"
)

// Builder builds converter implementations, and can decide if it can handle the given type.
type Builder interface {
	// Matches returns true, if the builder can create handle the given types.
	Matches(ctx *MethodContext, source, target *xtype.Type) bool
	// Build creates conversion source code for the given source and target type.
	Build(gen Generator, ctx *MethodContext, sourceID *xtype.JenID, source, target *xtype.Type) ([]jen.Code, *xtype.JenID, *Error)
}

// Generator checks all existing builders if they can create a conversion implementations for the given source and target type
// If no one Builder#Matches then, an error is returned.
type Generator interface {
	Build(
		ctx *MethodContext,
		sourceID *xtype.JenID,
		source, target *xtype.Type,
		errWrapper ErrorWrapper) ([]jen.Code, *xtype.JenID, *Error)

	CallMethod(
		ctx *MethodContext,
		method *method.Definition,
		sourceID *xtype.JenID,
		source, target *xtype.Type,
		errWrapper ErrorWrapper) ([]jen.Code, *xtype.JenID, *Error)
}

// MethodContext exposes information for the current method.
type MethodContext struct {
	*namer.Namer
	Conf              *config.Method
	FieldsTarget      string
	OutputPackagePath string
	Signature         xtype.Signature
	TargetType        *xtype.Type
	HasMethod         func(types.Type, types.Type) bool
	SeenNamed         map[string]struct{}

	TargetVar *jen.Statement
}

func (ctx *MethodContext) HasSeen(source *xtype.Type) bool {
	if !source.Named {
		return false
	}
	typeString := source.NamedType.String()
	_, ok := ctx.SeenNamed[typeString]
	return ok
}

func (ctx *MethodContext) MarkSeen(source *xtype.Type) {
	if !source.Named {
		return
	}
	typeString := source.NamedType.String()
	ctx.SeenNamed[typeString] = struct{}{}
}
func (ctx *MethodContext) HasAggregation(conf *config.Method, sourse *xtype.Type, target *xtype.Type) (place *xtype.Type, targetSlice *xtype.Type) {
	if conf.RawFieldSettings == nil {
		return nil, nil
	} else {
		if len(conf.RawFieldSettings) != 0 {
			split := strings.Split(conf.RawFieldSettings[0], " ")
			switch split[0] {
			case "agg":
				fmt.Println("agg", split[1], split[2])
				for i := 0; i < sourse.ListInner.StructType.NumFields(); i++ {
					// If the field name matches "name of place of aggregation", set the xtype.Type to the field information
					if split[1] == target.ListInner.StructType.Field(i).Name() {
						place = &xtype.Type{
							String: target.ListInner.StructType.Field(i).Name(),
							T:      target.ListInner.StructType.Field(i).Type(),
						}

					}
				}
				for i := 0; i < target.ListInner.StructType.NumFields(); i++ {
					// If the field name matches "[]", set the xtype.Type to the field information
					if split[2] == target.ListInner.StructType.Field(i).Name() {
						targetSlice = &xtype.Type{
							String: target.ListInner.StructType.Field(i).Name(),
							T:      target.ListInner.StructType.Field(i).Type(),
						}
					}
				}
				return place, targetSlice
			default:
				return nil, nil
			}
		}

	}
	return nil, nil
}
func (ctx *MethodContext) SetErrorTargetVar(m *jen.Statement) {
	if ctx.TargetVar == nil {
		ctx.TargetVar = m
	}
}

func (ctx *MethodContext) Field(target *xtype.Type, name string) *config.FieldMapping {
	if ctx.FieldsTarget != target.String {
		return emptyMapping
	}

	prop, ok := ctx.Conf.Fields[name]
	if !ok {
		return emptyMapping
	}
	return prop
}

func (ctx *MethodContext) DefinedFields(target *xtype.Type) map[string]struct{} {
	if ctx.FieldsTarget != target.String {
		return emptyFields
	}

	f := map[string]struct{}{}
	for name := range ctx.Conf.Fields {
		f[name] = struct{}{}
	}
	return f
}

var (
	emptyMapping *config.FieldMapping = &config.FieldMapping{}
	emptyFields                       = map[string]struct{}{}
)
