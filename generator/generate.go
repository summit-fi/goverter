package generator

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/emp1re/goverter-test/builder"
	"github.com/emp1re/goverter-test/config"
)

// Config the generate config.
type Config struct {
	WorkingDir      string
	BuildConstraint string
}

// BuildSteps that'll used for generation.
var BuildSteps = []builder.Builder{
	&builder.UseUnderlyingTypeMethods{},
	&builder.SkipCopy{},
	&builder.Agg{},
	&builder.ItemToListRule{},
	&builder.BasicTargetPointerRule{},
	&builder.Pointer{},
	&builder.SourcePointer{},
	&builder.TargetPointer{},
	&builder.Basic{},
	&builder.Struct{},
	&builder.List{},
	&builder.Map{},
}

// Generate generates a jen.File containing converters.
func Generate(converters []*config.Converter, c Config) (map[string][]byte, error) {
	manager := &fileManager{Files: map[string]*managedFile{}}
	for _, converter := range converters {
		jenFile, err := manager.Get(converter, c)
		if err != nil {
			return nil, err
		}

		if err := generateConverter(converter, c, jenFile); err != nil {
			return nil, err
		}
	}

	return manager.renderFiles()
}

func generateConverter(converter *config.Converter, c Config, f *jen.File) error {
	gen := setupGenerator(converter)
	cv := converter.Methods
	for _, method := range cv {

		for _, m := range method.RawFieldSettings {
			split := strings.Split(m, " ")
			if len(split) < 3 {
				return fmt.Errorf("Invalid settings for goverter:agg")
			}
			if split[0] == "agg" {
				if len(converter.Comments) > 0 {
					f.Comment(strings.Join(converter.Comments, "\n"))
				}
				f.Type().Id(converter.Name).Struct()

				if err := gen.buildMethods(f); err != nil {
					return err
				}
				return nil
			} else {
				continue
			}
		}
	}
	if err := validateMethods(gen.lookup); err != nil {
		return err
	}

	if len(converter.Comments) > 0 {
		f.Comment(strings.Join(converter.Comments, "\n"))
	}
	f.Type().Id(converter.Name).Struct()

	if err := gen.buildMethods(f); err != nil {
		return err
	}
	return nil
}
