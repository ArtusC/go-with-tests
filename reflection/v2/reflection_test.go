//go:build unit
// +build unit

package reflection_test

import (
	"reflect"
	"testing"

	re "github.com/ArtusC/go-with-tests/reflection/v2"
)

func TestWalk(t *testing.T) {

	type testCase struct {
		CaseName      string
		Input         interface{}
		ExpectedCalls []string
	}

	type Input_1 struct {
		Name string
	}

	type Input_2 struct {
		Name string
		City string
	}

	cases := []testCase{
		{
			CaseName:      "struct with one string field",
			Input:         Input_1{Name: "Artus"},
			ExpectedCalls: []string{"Artus"},
		},
		{
			CaseName:      "struct with two string fields",
			Input:         Input_2{Name: "Artus", City: "Floripa"},
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
