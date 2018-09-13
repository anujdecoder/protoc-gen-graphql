package templates

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/srikrsna/protoc-gen-graphql/graphql"
	"strings"
	"text/template"
)

var T = template.New("File").Funcs(template.FuncMap{
	"trimSpace": strings.TrimSpace,
	"getKey": func(m pgs.Message) string {
		id := false
		customKey := ""
		for _, v := range m.NonOneOfFields() {
			// TODO: Check using extension for custom keys
			// TODO: Also check if it's go scalar as thunder uses go scalars

			// Checking for most common identifier 'Id'
			if v.Name().LowerCamelCase() == "id" {
				id = true
			}
		}

		if customKey != "" {
			return customKey
		}

		if id {
			return "id"
		}

		return ""
	},
	"getMethodType": func(m pgs.Method) (*struct {
		Type       string
		MethodName string
	}, error) {
		var o graphql.MethodOptions
		if ok, err := m.Extension(graphql.E_GraphqlMethods, &o); err != nil || !ok {
			return nil, err
		}

		switch v := o.Type.(type) {
		case *graphql.MethodOptions_Query:
			return &struct {
				Type       string
				MethodName string
			}{
				Type:       "Query",
				MethodName: v.Query,
			}, nil
		case *graphql.MethodOptions_Mutation:
			return &struct {
				Type       string
				MethodName string
			}{
				Type:       "Mutation",
				MethodName: v.Mutation,
			}, nil
		}

		return nil, nil
	},
	"getImports": func(f pgs.File) ([]pgs.Package, error) {
		additionalImports := map[string]pgs.Package{}

		for _, v := range f.Services() {
			for _, v := range v.Methods() {
				var options graphql.MethodOptions

				ok, err := v.Extension(graphql.E_GraphqlMethods, &options)
				if err != nil {
					return nil, err
				}

				if !ok {
					continue
				}

				if !v.Input().BuildTarget() {
					additionalImports[v.Input().Package().GoName().String()] = v.Input().Package()
				}

				if !v.Output().BuildTarget() {
					additionalImports[v.Output().Package().GoName().String()] = v.Output().Package()
				}
			}
		}

		var imports []pgs.Package

		for _, v := range additionalImports {
			imports = append(imports, v)
		}

		return imports, nil
	},
})
