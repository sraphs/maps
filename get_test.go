package maps

import (
	"reflect"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	m := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"nil":          nil,
				"bool":         true,
				"int_key":      1000,
				"float_key":    1000.1,
				"string_key":   "string_value",
				"duration_key": 10000,
			},
		},
	}

	tests := []struct {
		name   string
		path   string
		ok     bool
		expect interface{}
	}{
		{
			name:   "nil",
			path:   "a.b.nil",
			ok:     true,
			expect: nil,
		},
		{
			name:   "bool",
			path:   "a.b.bool",
			ok:     true,
			expect: true,
		},
		{
			name:   "int_key",
			path:   "a.b.int_key",
			ok:     true,
			expect: int64(1000),
		},
		{
			name:   "float_key",
			path:   "a.b.float_key",
			ok:     true,
			expect: float64(1000.1),
		},
		{
			name:   "string_key",
			path:   "a.b.string_key",
			ok:     true,
			expect: "string_value",
		},
		{
			name:   "duration_key",
			path:   "a.b.duration_key",
			ok:     true,
			expect: time.Duration(10000),
		},
		{
			name:   "not found",
			path:   "a.b.Z",
			ok:     false,
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Get(m, tt.path)

			var err error
			var actual interface{}

			switch tt.expect.(type) {
			case int64:
				if actual, err = v.Int(); err == nil {
					if !reflect.DeepEqual(tt.expect, actual) {
						t.Errorf("expect %v, actual %v", tt.expect, actual)
					}
				}
			case string:
				if actual, err = v.String(); err == nil {
					if !reflect.DeepEqual(tt.expect, actual) {
						t.Errorf(`expect %v, actual %v`, tt.expect, actual)
					}
				}
			case bool:
				if actual, err = v.Bool(); err == nil {
					if !reflect.DeepEqual(tt.expect, actual) {
						t.Errorf(`expect %v, actual %v`, tt.expect, actual)
					}
				}
			case float64:
				if actual, err = v.Float(); err == nil {
					if !reflect.DeepEqual(tt.expect, actual) {
						t.Errorf(`expect %v, actual %v`, tt.expect, actual)
					}
				}
			case time.Duration:
				if actual, err = v.Duration(); err == nil {
					if !reflect.DeepEqual(tt.expect, actual) {
						t.Errorf(`expect %v, actual %v`, tt.expect, actual)
					}
				}
			default:
				actual = v.Load()
				if !reflect.DeepEqual(tt.expect, actual) {
					t.Logf("\nexpect: %#v\nactural: %#v", tt.expect, actual)
					t.Fail()
				}
			}
		})
	}
}
