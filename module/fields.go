package module

import (
	"github.com/lyft/protoc-gen-star"
	gotag "github.com/srikrsna/protoc-gen-gotag/module"
	"strconv"
	"strings"
)

import "github.com/fatih/structtag"

type TagCase int

const (
	LowerCamelCase TagCase = iota
	UpperCamelCase
	LowerSnakeCase
	UpperSnakeCase
	ScreamingSnakeCase
	LowerDotCase
	UpperDotCase
)

// MakeTags extracts all tags in the loweCamelCaseFormat
func MakeTags(d pgs.DebuggerCommon, n pgs.Node, tagCase TagCase, wk map[string]WellKnown) (gotag.StructTags, error) {
	v := newFieldVisitor(tagCase, d, wk)

	if err := pgs.Walk(v, n); err != nil {
		return nil, err
	}

	return v.tags, nil
}

type tagExtractor struct {
	pgs.Visitor
	pgs.DebuggerCommon

	tags gotag.StructTags

	mod, first pgs.NameTransformer
	sep        string

	wk map[string]WellKnown
}

func newFieldVisitor(tagCase TagCase, d pgs.DebuggerCommon, wk map[string]WellKnown) *tagExtractor {

	v := tagExtractor{tags: map[string]map[string]*structtag.Tags{}, DebuggerCommon: d, wk: wk}

	switch tagCase {
	case LowerCamelCase:
		v.mod, v.first, v.sep = strings.Title, strings.ToLower, ""
	case UpperCamelCase:
		v.mod, v.first, v.sep = strings.Title, strings.Title, ""
	case LowerSnakeCase:
		v.mod, v.first, v.sep = strings.ToLower, strings.ToLower, "_"
	case UpperSnakeCase:
		v.mod, v.first, v.sep = strings.Title, strings.Title, "_"
	case ScreamingSnakeCase:
		v.mod, v.first, v.sep = strings.ToUpper, strings.ToUpper, "_"
	case LowerDotCase:
		v.mod, v.first, v.sep = strings.ToLower, strings.ToLower, "."
	case UpperDotCase:
		v.mod, v.first, v.sep = strings.Title, strings.Title, "."
	}

	v.Visitor = pgs.PassThroughVisitor(&v)
	return &v
}

func (v *tagExtractor) VisitOneOf(o pgs.OneOf) (pgs.Visitor, error) {
	msgName := o.Message().Name().PGGUpperCamelCase().String()
	if v.tags[msgName] == nil {
		v.tags[msgName] = map[string]*structtag.Tags{}
	}

	tags, err := structtag.Parse(`graphql:"-"`)
	if err != nil {
		return nil, err
	}

	v.tags[msgName][o.Name().PGGUpperCamelCase().String()] = tags

	return v, nil
}

func (v *tagExtractor) VisitField(f pgs.Field) (pgs.Visitor, error) {
	msgName := f.Message().Name().PGGUpperCamelCase().String()

	if f.InOneOf() {
		msgName = f.Message().Name().PGGUpperCamelCase().String() + "_" + f.Name().PGGUpperCamelCase().String()
	}

	if v.tags[msgName] == nil {
		v.tags[msgName] = map[string]*structtag.Tags{}
	}

	tags, err := structtag.Parse("graphql:"+strconv.Quote(f.Name().Transform(v.mod, v.first, v.sep).String()))
	if err != nil {
		return nil, err
	}

	if f.Descriptor().TypeName != nil {
		if _, ok := v.wk[*f.Descriptor().TypeName]; ok {
			tags, err = structtag.Parse(`graphql:"-"`)
			if err != nil {
				return nil, err
			}
		}
	}

	v.tags[msgName][f.Name().PGGUpperCamelCase().String()] = tags

	return v, nil
}

func (v *tagExtractor) VisitMessage(f pgs.Message) (pgs.Visitor, error) {
	name := f.Name().PGGUpperCamelCase().String()
	if v.tags[name] == nil {
		v.tags[name] = map[string]*structtag.Tags{}
	}

	return v, nil
}
