package module // import "github.com/srikrsna/protoc-gen-graphql/module"

import (
	"github.com/fatih/structtag"
	"github.com/lyft/protoc-gen-star"
	gtag "github.com/srikrsna/protoc-gen-gotag/module"
	"github.com/srikrsna/protoc-gen-graphql/templates"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

type mod struct {
	*pgs.ModuleBase
}

func New() pgs.Module {
	return &mod{&pgs.ModuleBase{}}
}

func (m mod) Name() string {
	return "graphql"
}

func (m mod) Execute(target pgs.Package, packages map[string]pgs.Package) []pgs.Artifact {

	xxxTags, err := structtag.Parse(`graphql:"-"`)
	m.CheckErr(err)

	wks := GTypes()

	for _, f := range target.Files() {
		tags, err := MakeTags(m, f, LowerCamelCase, wks)
		m.CheckErr(err)

		tags.AddTagsToXXXFields(xxxTags)

		gon := f.OutputPath().SetExt(".go").String()

		fs := token.NewFileSet()
		fn, err := parser.ParseFile(fs, gon, nil, parser.ParseComments)
		if err != nil {
			m.Failf("unable to parse file into ast: %v", err)
		}

		m.CheckErr(gtag.Retag(fn, tags))

		var buf strings.Builder
		m.CheckErr(printer.Fprint(&buf, fs, fn))

		m.OverwriteGeneratorFile(gon, buf.String())

		m.AddGeneratorTemplateFile(
			f.OutputPath().SetExt(".gl.go").String(),
			templates.T,
			f,
		)
	}

	return m.Artifacts()
}
