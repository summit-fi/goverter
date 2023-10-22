package pkgload

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func New(workDir string) *PackageLoader {
	return &PackageLoader{
		cache:      map[string][]*packages.Package{},
		workingDir: workDir,
	}
}

type PackageLoader struct {
	workingDir string
	cache      map[string][]*packages.Package
}

// loadPackages is used to load extend packages, with caching support.
func (g *PackageLoader) Load(pkgPath string) ([]*packages.Package, error) {
	if pkgs, ok := g.cache[pkgPath]; ok {
		return pkgs, nil
	}

	packagesCfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  g.workingDir,
	}
	pkgs, err := packages.Load(packagesCfg, pkgPath)
	if err != nil {
		// This happens rare, and only if somebody uses advanced package pattern query in a wrong way.
		// The cause (err) usually has enough details to troubleshoot this issue.
		return nil, fmt.Errorf("packages load failed on %q: %s", pkgPath, err)
	}
	// we need at least one valid package with no errors reported during its load
	var hasValidPackage bool
	var firstErr error
	for _, pkg := range pkgs {
		if len(pkg.Errors) == 0 {
			hasValidPackage = true
			break
		}
		if firstErr == nil {
			firstErr = pkg.Errors[0]
		}
	}
	if !hasValidPackage {
		// no valid package detected, report this as an error
		if firstErr == nil {
			// Most of the time, if packages.Load fails to load a package, it will still return
			// this pkg with PkgPath = input and pkg.Errors with at least one error indicating what
			// happened. However, if somebody uses advanced package pattern like file=/path, and
			// this path does not exist, then packages.Load does not fail, yet it also returns no
			// packages. We need to fail this case, and we cannot suggest using blank import for such
			// cases, thus using a generic error.
			return nil, fmt.Errorf("no packages were loaded for %q, make sure it is a valid golang package", pkgPath)
		}
		// Packages.Load may need local go module's help to load external packages, the best way
		// to help is to load same package directly into converter's module using a blank import.
		return nil, fmt.Errorf("failed to load package %q, try adding a blank import for it: %s", pkgPath, firstErr)
	}

	g.cache[pkgPath] = pkgs
	return pkgs, nil
}
