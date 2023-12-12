//go:build unit
// +build unit

package reflection_test

import (
	"reflect"
	"testing"

	re "github.com/ArtusC/go-with-tests/reflection/v6"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City   string
	Age    int
	Weight float32
}

func TestWalk(t *testing.T) {

	type testCase struct {
		CaseName      string
		Input         interface{}
		ExpectedCalls []string
	}

	cases := []testCase{
		{
			CaseName: "struct with one string field",
			Input: struct {
				Name string
			}{Name: "Artus"},
			ExpectedCalls: []string{"Artus"},
		},
		{
			CaseName: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{Name: "Artus", City: "Floripa"},
			ExpectedCalls: []string{"Artus", "Floripa"},
		},
		{
			CaseName: "struct with non string fields",
			Input: struct {
				Name   string
				City   string
				Age    int
				Weight float32
			}{Name: "Artus", City: "Floripa", Age: 31, Weight: 66.6},
			ExpectedCalls: []string{"Artus", "Floripa", "31", "66.60"},
		},
		{
			CaseName: "nested fields",
			Input: Person{
				Name: "Artus",
				Profile: Profile{
					City:   "Floripa",
					Age:    31,
					Weight: 66.6,
				}},
			ExpectedCalls: []string{"Artus", "Floripa", "31", "66.60"},
		},
		{
			CaseName: "pointers",
			Input: &Person{
				Name: "Artus",
				Profile: Profile{
					City:   "Floripa",
					Age:    31,
					Weight: 66.6,
				}},
			ExpectedCalls: []string{"Artus", "Floripa", "31", "66.60"},
		},
		{
			CaseName: "slices",
			Input: []Profile{
				{
					City:   "Floripa",
					Age:    31,
					Weight: 66.6,
				},
				{
					City:   "Urubici",
					Age:    32,
					Weight: 86.6,
				},
			},
			ExpectedCalls: []string{"Floripa", "31", "66.60", "Urubici", "32", "86.60"},
		},
		{
			CaseName: "arrays",
			Input: [2]Profile{
				{
					City:   "Floripa",
					Age:    31,
					Weight: 66.6,
				},
				{
					City:   "Urubici",
					Age:    32,
					Weight: 86.6,
				},
			},
			ExpectedCalls: []string{"Floripa", "31", "66.60", "Urubici", "32", "86.60"},
		},
	}

	for _, test := range cases {
		t.Run(test.CaseName, func(t *testing.T) {
			var got []string
			re.Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
