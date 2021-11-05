package reflection

import (
	"reflect"
	"testing"
)

type (
	Person struct {
		Name    string
		Profile Profile
	}
	Profile struct {
		Age  int
		City string
	}
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{{
		"Struct with one string field",
		struct {
			Name string
		}{"Chris"},
		[]string{"Chris"},
	},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 39},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{"Chris", Profile{39, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to thing",
			&Person{"Chris", Profile{39, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{39, "London"},
				{19, "Paris"},
			},
			[]string{"London", "Paris"},
		},
		{"Arrays",
			[2]Profile{
				{39, "London"},
				{19, "Paris"},
			},
			[]string{"London", "Paris"},
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
			"City": "London",
			"Name": "Paris",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "London")
		assertContains(t, got, "Paris")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{19, "Paris"}
			aChannel <- Profile{39, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Paris", "London"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{19, "Paris"}, Profile{39, "London"}
		}

		var got []string
		want := []string{"Paris", "London"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	var contains bool
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v, to cotain %q, but it didn't", haystack, needle)
	}

}
