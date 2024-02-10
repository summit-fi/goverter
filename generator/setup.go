package generator

import (
	"github.com/emp1re/goverter-test/config"
	"github.com/emp1re/goverter-test/method"
	"github.com/emp1re/goverter-test/namer"
	"github.com/emp1re/goverter-test/xtype"
)

func setupGenerator(converter *config.Converter) *generator {
	extend := map[xtype.Signature]*method.Definition{}
	for _, def := range converter.Extend {
		extend[def.Signature()] = def
	}

	lookup := map[xtype.Signature]*generatedMethod{}
	for _, method := range converter.Methods {
		lookup[method.Definition.Signature()] = &generatedMethod{
			Method:   method,
			Dirty:    false,
			Explicit: true,
		}
	}

	gen := generator{
		namer:  namer.New(),
		conf:   converter,
		lookup: lookup,
		extend: extend,
	}

	return &gen
}
