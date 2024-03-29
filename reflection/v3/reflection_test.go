//go:build unit
// +build unit

package reflection_test

import (
	"reflect"
	"testing"

	re "github.com/ArtusC/go-with-tests/reflection/v3"
)

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
			CaseName: "struct with non string field",
			Input: struct {
				Name string
				City string
				Age  int
			}{Name: "Artus", City: "Floripa", Age: 31},
			ExpectedCalls: []string{"Artus", "Floripa"},
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
