package module

import "github.com/lyft/protoc-gen-star"

type WellKnown struct {
	FullyQualifiedIdentifier string
	ReplaceWith              string
	IsGType                  bool
	ConvertFunc              string
}

func GTypes() map[string]WellKnown {
	wks := []WellKnown{
		{
			".google.protobuf.Any", "*gtypes.Any", true, "gtypes.AnyToScalar",
		},
		{
			".google.protobuf.Duration", "string", true, "gtypes.DurationToScalar",
		},
		{
			".google.protobuf.Timestamp", "time.Time", true, "gtypes.TimestampToScalar",
		},
		{
			".google.protobuf.Empty", "", true, "",
		},
	}

	res := map[string]WellKnown{}

	for _, v := range wks {
		res[v.FullyQualifiedIdentifier] = v
	}

	return res
}

func getWellKnownFunc(wks map[string]WellKnown) func(m pgs.Message) interface{} {
	f := func(m pgs.Message) (interface{}) {
		type wk struct {
			pgs.Field
			WellKnown
		}

		var res []wk

		for _, f := range m.NonOneOfFields() {

			name := f.Descriptor().TypeName

			if name == nil {
				continue
			}

			if v, ok := wks[*name]; ok {
				res = append(res, wk{
					Field:     f,
					WellKnown: v,
				})
			}
		}

		return res
	}

	return f
}
