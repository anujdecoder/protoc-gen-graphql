package gtypes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"strings"
	"time"
)

var Schema = schemabuilder.NewSchema()

func init() {
	RegisterWellKnownTypes(Schema)
}

func RegisterWellKnownTypes(s *schemabuilder.Schema) {
	RegisterAny(s)
}

func RegisterAny(s *schemabuilder.Schema) {
	obj := s.Object("Google.Protobuf.Any", Any{})
	obj.FieldFunc("typeUrl", func(in *Any) string {
		return in.in.TypeUrl
	})
	obj.FieldFunc("value", func(in *Any) []byte {
		return in.in.Value
	})
}

type Any struct {
	in *any.Any `graphql:"-"`
}

func AnyToScalar(a *any.Any) (*Any, error) {
	return &Any{in: a}, nil
}

func AnyFromScalar(a *Any) (*any.Any, error) {
	return a.in, nil
}

func TimestampToScalar(t *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(t)
}

func TimestampFromScalar(t time.Time) (*timestamp.Timestamp, error) {
	return ptypes.TimestampProto(t)
}

const secondInNanos = int64(time.Second / time.Nanosecond)

func DurationToScalar(d duration.Duration) (string, error) {
	s, ns := d.Seconds, int64(d.Nanos)
	if ns <= -secondInNanos || ns >= secondInNanos {
		return "", fmt.Errorf("ns out of range (%v, %v)", -secondInNanos, secondInNanos)
	}
	if (s > 0 && ns < 0) || (s < 0 && ns > 0) {
		return "", errors.New("signs of seconds and nanos do not match")
	}
	if s < 0 {
		ns = -ns
	}
	var out strings.Builder

	x := fmt.Sprintf("%d.%09d", s, ns)
	x = strings.TrimSuffix(x, "000")
	x = strings.TrimSuffix(x, "000")
	x = strings.TrimSuffix(x, ".000")

	out.WriteRune('"')
	out.WriteString(x)
	out.WriteString(`s"`)

	return out.String(), nil
}

func DurationFromScalar(in string) (*duration.Duration, error) {
	unq, err := unquote(in)
	if err != nil {
		return nil, err
	}

	d, err := time.ParseDuration(unq)
	if err != nil {
		return nil, fmt.Errorf("bad Duration: %v", err)
	}

	ns := d.Nanoseconds()
	s := ns / 1e9
	ns %= 1e9

	return &duration.Duration{
		Seconds: s,
		Nanos:   int32(ns),
	}, nil
}

func unquote(s string) (string, error) {
	var ret string
	err := json.Unmarshal([]byte(s), &ret)
	return ret, err
}
