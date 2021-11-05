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
		{"Map",
			map[string]string{
				"City": "London",
				"Name": "Paris",
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

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %s, want %s", got[0], expected)
	}

}
