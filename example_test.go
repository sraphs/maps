package maps_test

import (
	"fmt"

	"github.com/sraphs/maps"
)

func Example() {
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

	n := maps.Get(m, "a.b.nil").Load()
	b, _ := maps.Get(m, "a.b.bool").Bool()
	i, _ := maps.Get(m, "a.b.int_key").Int()
	f, _ := maps.Get(m, "a.b.float_key").Float()
	s, _ := maps.Get(m, "a.b.string_key").String()
	d, _ := maps.Get(m, "a.b.duration_key").Duration()

	fmt.Println(n)
	fmt.Println(b)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(d)

	//Output:
	// <nil>
	// true
	// 1000
	// 1000.1
	// string_value
	// 10µs
}
