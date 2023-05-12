package database

import (
	"fmt"
	"testing"
)

func TestGetDataForUser(t *testing.T) {
	baseUrl := "../"
	cases := []struct {
		username string
		length   int
	}{
		{"michael", 2},
		{"rachel", 1},
		{"john", 0},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("running test for user %s", testCase.username), func(t *testing.T) {
			got, err := GetDataForUser(testCase.username, baseUrl)
			if err != nil {
				t.Error("got and error: ", err)
			}
			if len(got) != testCase.length {
				t.Errorf("expected %d, but got %d", testCase.length, len(got))
			}

		})
	}

}
