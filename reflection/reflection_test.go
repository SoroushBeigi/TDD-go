package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Soroush"},
			[]string{"Soroush"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Soroush", "Tehran"},
			[]string{"Soroush", "Tehran"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Soroush", 21},
			[]string{"Soroush"},
		},
		{
			"nested fields",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{
				"Soroush", struct {
					Age  int
					City string
				}{21, "Tehran"}},
			[]string{"Soroush", "Tehran"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
