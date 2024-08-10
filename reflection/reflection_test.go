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
			[2]Profile{
				{21, "Tehran"},
				{22, "Berlin"},
			},
			[]string{"Tehran", "Berlin"},
		},
		{
			"maps",
			map[string]string{
				"City": "Tehran",
				"Age":  "21",
			},
			[]string{"Tehran", "21"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{21, "Tehran"}
			aChannel <- Profile{30, "Paris"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Tehran", "Paris"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with func", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{21, "Tehran"}, Profile{55, "Amsterdam"}
		}
		var got []string
		want := []string{"Tehran", "Amsterdam"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, list []string, item string) {
	t.Helper()
	contains := false
	for _, x := range list {
		if x == item {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q", list, item)
	}
}
