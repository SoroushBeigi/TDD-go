package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
			Person{
				"Soroush",
				Profile{21, "Tehran"},
			},
			[]string{"Soroush", "Tehran"},
		},
		{
			"pointers",
			&Person{
				"Soroush",
				Profile{21, "Tehran"},
			},
			[]string{"Soroush", "Tehran"},
		},
		{
			"slices",
			[]Profile{
				{21, "Tehran"},
				{54, "London"},
			},
			[]string{"Tehran", "London"},
		},
		{
			"arrays",
			[2]Profile {
				{21, "Tehran"},
				{22, "Berlin"},
			},
			[]string{"Tehran", "Berlin"},
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
