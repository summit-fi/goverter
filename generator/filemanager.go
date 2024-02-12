package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/emp1re/goverter-test/config"
)

type fileManager struct {
	Files map[string]*managedFile
}

type managedFile struct {
	PackageID string
	Initial   *config.Converter
	Content   *jen.File
}

func (m *fileManager) Get(conv *config.Converter, cfg Config) (*jen.File, error) {
	output := getOutputDir(conv, cfg.WorkingDir)

	f, ok := m.Files[output]
	if !ok {
		f = &managedFile{
			PackageID: conv.PackageID(),
			Initial:   conv,
		}

		if conv.OutputPackageName == "" {
			f.Content = jen.NewFilePath(conv.OutputPackagePath)
		} else {
			f.Content = jen.NewFilePathName(conv.OutputPackagePath, conv.OutputPackageName)
		}

		f.Content.HeaderComment("// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.")
		if cfg.BuildConstraint != "" {
			f.Content.HeaderComment("//go:build " + cfg.BuildConstraint)
		}
		m.Files[output] = f
	}

	if f.PackageID != conv.PackageID() {
		return nil, fmt.Errorf("Error creating converters\n    %s\n    %s\nand\n    %s\n    %s\n\nCannot use different packages\n    %s\n    %s\nin the same output file:\n    %s",
			conv.FileSource, conv.Type, f.Initial.FileSource, f.Initial.Type, conv.PackageID(), f.Initial.PackageID(), output)
	}

	return f.Content, nil
}

func (m *fileManager) renderFiles() (map[string][]byte, error) {
	result := map[string][]byte{}
	for name, f := range m.Files {
		var buf bytes.Buffer
		if err := f.Content.Render(&buf); err != nil {
			return result, err
		}
		result[name] = buf.Bytes()
	}
	return result, nil
}

func getOutputDir(c *config.Converter, cwd string) string {
	if strings.HasPrefix(c.OutputFile, "@cwd/") {
		return filepath.Join(cwd, strings.TrimPrefix(c.OutputFile, "@cwd/"))
	}

	if filepath.IsAbs(c.OutputFile) {
		return c.OutputFile
	}

	return filepath.Join(filepath.Dir(c.FileSource), c.OutputFile)
}
