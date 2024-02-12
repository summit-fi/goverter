package builder

import (
	"go/types"

	"github.com/dave/jennifer/jen"
	"github.com/emp1re/goverter-test/config"
	"github.com/emp1re/goverter-test/method"
	"github.com/emp1re/goverter-test/namer"
	"github.com/emp1re/goverter-test/xtype"
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
